package list

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/api"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Incubator/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
	"net/url"
	"strconv"
)

func NewCmdScheduledViewsList() *cobra.Command {
	var limit int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all scheduled views in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			listScheduledViews(limit)
		},
	}
	cmd.Flags().IntVar(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	return cmd
}

func listScheduledViews(limit int) {
	var scheduledViewsResponse api.GetScheduledViews
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/scheduledViews"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &scheduledViewsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	scheduledViewsResponseJson, err := json.MarshalIndent(scheduledViewsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(scheduledViewsResponseJson))
	}
}
