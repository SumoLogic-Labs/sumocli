package start

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
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
	State           liveTailSessionInfo `json:"state"`
	Messages        []string            `json:"messages"`
	KeyedErrors     []string            `json:keyedErrors"`
	Error           bool                `json:"error"`
	ErrorMessage    string              `json:"errorMessage"`
	ErrorInstanceId string              `json:"errorInstanceId"`
	ErrorKey        string              `json:"errorKey"`
	ResultTraitLog  string              `json:"resultTraitLog"`
}

type liveTailSessionInfo struct {
	TailId        int      `json:"tailId"`
	CurrentOffset int      `json:"currentOffset"`
	IsStopped     bool     `json:"isStopped"`
	UserMessages  []string `json:"userMessages"`
	Stopped       bool     `json:"stopped"`
}

func StartLiveTailCmd() *cobra.Command {
	var (
		filter string
		tailId string
	)
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a live tail session",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			startLiveTailSession(log)
		},
	}

	cmd.Flags().StringVar(&filter, "filter", "", "")
	cmd.Flags().StringVar(&tailId, "tailId", "", "Test argument")
	return cmd
}

func IsEmpty(response liveTailSessionInfo) bool {
	return reflect.DeepEqual(liveTailSessionInfo{}, response)
}

func createLiveTailSession(log zerolog.Logger) string {
	var session createLiveTailSessionResponse
	accessId, accessKey, endpoint := login.ReadAccessKeys()
	requestBodySchema := &CreateLiveTailSessionRequest{
		IsCLI:  true,
		Filter: "_sourceCategory=test",
	}
	requestBody, _ := json.Marshal(requestBodySchema)
	client, request := factory.NewLiveTailHttpRequest("POST", endpoint+"v1/livetail/session", requestBody)
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
	}
	return session.Id
}

func startLiveTailSession(log zerolog.Logger) {
	sessionId := createLiveTailSession(log)
	accessId, accessKey, endpoint := login.ReadAccessKeys()
	offset := 0
	fmt.Println(sessionId)

	for true {
		latestLiveTailResultsUrl := endpoint + "v1/livetail/session/" + sessionId + "/latest/" + strconv.Itoa(offset)
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
			for i, userMessage := range tailSession.State.UserMessages {
				fmt.Println(userMessage)
				i++
			}
		}

		fmt.Println(tailSession.State.TailId)
		fmt.Println(tailSession.State.CurrentOffset)
		fmt.Println(tailSession.State.IsStopped)
		fmt.Println(tailSession.State.Stopped)
		fmt.Println(tailSession.State.UserMessages)
	}
}
