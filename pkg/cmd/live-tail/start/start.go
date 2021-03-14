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
	"time"
)

type CreateLiveTailSessionRequest struct {
	IsCLI  bool   `json:"isCLI""`
	Filter string `json:"filter,omitempty"`
}

type CreateLiveTailSessionResponse struct {
	Id              string   `json:"id"`
	StartTime       int      `json:"startTime"`
	KeyedErrors     []string `json:"keyedErrors"`
	ErrorMessage    string   `json:"errorMessage"`
	ErrorInstanceId string   `json:"errorInstanceId"`
	ErrorKey        string   `json:"errorKey"`
	ResultTraitLog  string   `json:"resultTraitLog"`
}

func StartLiveTailCmd() *cobra.Command {
	var (
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

	cmd.Flags().StringVar(&tailId, "tailId", "", "Test argument")
	return cmd
}

func createLiveTailSession(log zerolog.Logger) string {
	var session CreateLiveTailSessionResponse
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
		latestLiveTailResultsUrl := endpoint + "v1/livetail/session/" + sessionId + "/latest/" + string(offset)
		time.Sleep(2 * time.Second)

		func() {
			fmt.Println(latestLiveTailResultsUrl)
			fmt.Println(accessId)
			fmt.Println(accessKey)
			client, request := factory.StartLiveTailHttpRequest("POST", latestLiveTailResultsUrl)
			request.SetBasicAuth(accessId, accessKey)
			_, _ = client.Do(request)
			/*
				client, request := factory.StartLiveTailHttpRequest("POST", latestLiveTailResultsUrl)
				request.SetBasicAuth(accessId, accessKey)
				response, err := client.Do(request)
				if err != nil {
					log.Error().Err(err).Msg("failed to make http request to " + endpoint)
				}
				defer response.Body.Close()
				responseBody, err := io.ReadAll(response.Body)
				return responseBody

			*/
		}()
	}
}
