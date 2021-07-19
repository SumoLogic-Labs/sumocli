package start

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/internal/authentication"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"reflect"
	"strconv"
	"time"
)

type CreateLiveTailSessionRequest struct {
	IsCLI  bool   `json:"isCLI""`
	Filter string `json:"filter,omitempty"`
}

type createLiveTailSessionResponse struct {
	Id              string   `json:"id"`
	StartTime       int      `json:"startTime"`
	KeyedErrors     []string `json:"keyedErrors"`
	ErrorMessage    string   `json:"errorMessage"`
	ErrorInstanceId string   `json:"errorInstanceId"`
	ErrorKey        string   `json:"errorKey"`
	ResultTraitLog  string   `json:"resultTraitLog"`
}

type liveTailSessionResponse struct {
	State           liveTailSessionInfo              `json:"state"`
	Messages        []liveTailSessionMessageResponse `json:"messages"`
	KeyedErrors     []string                         `json:keyedErrors"`
	Error           bool                             `json:"error"`
	ErrorMessage    string                           `json:"errorMessage"`
	ErrorInstanceId string                           `json:"errorInstanceId"`
	ErrorKey        string                           `json:"errorKey"`
	ResultTraitLog  string                           `json:"resultTraitLog"`
}

type liveTailSessionMessageResponse struct {
	Offset         int    `json:"offset"`
	ReceiptTime    int    `json:"receiptTime"`
	MessageTime    int    `json:"messageTime"`
	Payload        string `json:"payload"`
	SourceName     string `json:"sourceName"`
	SourceHost     string `json:"sourceHost"`
	SourceCategory string `json:"sourceCategory"`
}

type liveTailSessionInfo struct {
	TailId        int      `json:"tailId"`
	CurrentOffset int      `json:"currentOffset"`
	IsStopped     bool     `json:"isStopped"`
	UserMessages  []string `json:"userMessages"`
	Stopped       bool     `json:"stopped"`
}

func StartLiveTailCmd() *cobra.Command {
	var filter string
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a live tail session",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			startLiveTailSession(filter, log)
		},
	}

	cmd.Flags().StringVar(&filter, "filter", "", "")
	cmd.MarkFlagRequired(filter)
	return cmd
}

func IsEmpty(response liveTailSessionInfo) bool {
	return reflect.DeepEqual(liveTailSessionInfo{}, response)
}

func createLiveTailSession(filter string, log zerolog.Logger) string {
	var session createLiveTailSessionResponse
	accessId, accessKey, endpoint := authentication.ReadAccessKeys()
	requestBodySchema := &CreateLiveTailSessionRequest{
		IsCLI:  true,
		Filter: "_sourceCategory=ubuntu/syslog",
	}
	requestBody, _ := json.Marshal(requestBodySchema)
	client, request := factory.NewLiveTailHttpRequest("POST", endpoint+"/v1/livetail/session", requestBody)
	request.SetBasicAuth(accessId, accessKey)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + endpoint)
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}
	err = json.Unmarshal(responseBody, &session)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
		fmt.Println(string(responseBody))
	}
	return session.Id
}

func startLiveTailSession(filter string, log zerolog.Logger) {
	sessionId := createLiveTailSession(filter, log)
	accessId, accessKey, endpoint := authentication.ReadAccessKeys()
	fmt.Println("### Starting Live Tail Session ###")
	offset := 0

	for true {
		latestLiveTailResultsUrl := endpoint + "/v1/livetail/session/" + sessionId + "/latest/" + strconv.Itoa(offset)
		time.Sleep(2 * time.Second)
		tailSession := func() liveTailSessionResponse {
			var tailSession liveTailSessionResponse
			client, request := factory.StartLiveTailHttpRequest("GET", latestLiveTailResultsUrl)
			request.SetBasicAuth(accessId, accessKey)
			response, err := client.Do(request)
			if err != nil {
				log.Error().Err(err).Msg("failed to make http request to " + endpoint)
			}
			defer response.Body.Close()
			responseBody, err := io.ReadAll(response.Body)
			if err != nil {
				log.Error().Err(err).Msg("error reading response body from request")
			}
			err = json.Unmarshal(responseBody, &tailSession)
			if err != nil {
				log.Error().Err(err).Msg("error unmarshalling response body")
			}
			return tailSession
		}()
		if IsEmpty(tailSession.State) == false {
			offset = tailSession.State.CurrentOffset + 1
			messages := tailSession.Messages

			for i, userMessage := range tailSession.State.UserMessages {
				fmt.Println(userMessage)
				i++
			}

			for i, message := range messages {
				fmt.Println(message.Payload)
				i++
			}
		}
	}
}
