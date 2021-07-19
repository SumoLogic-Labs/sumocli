package list

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"net/url"
	"strconv"
)

func NewCmdHealthEventsList() *cobra.Command {
	var limit int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all the unresolved health events in your account.",
		Run: func(cmd *cobra.Command, args []string) {
			listHealthEvents(limit)
		},
	}
	cmd.Flags().IntVar(&limit, "limit", 100, "Specify the number of health events to return")
	return cmd
}

func listHealthEvents(limit int) {
	var healthEventsResponse api.ListHealthEvent
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/healthEvents"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &healthEventsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	healthEventsResponseJson, err := json.MarshalIndent(healthEventsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(healthEventsResponseJson))
	}
}
