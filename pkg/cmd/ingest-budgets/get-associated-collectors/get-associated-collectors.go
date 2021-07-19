package get_associated_collectors

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

func NewCmdIngestBudgetsGetAssociatedCollectors() *cobra.Command {
	var (
		id    string
		limit int
	)

	cmd := &cobra.Command{
		Use:   "get-associated-collectors",
		Short: "Get a list of Collectors assigned to an ingest budget.",
		Run: func(cmd *cobra.Command, args []string) {
			getAssociatedCollectors(id, limit)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.Flags().IntVar(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getAssociatedCollectors(id string, limit int) {
	var associatedCollectorsResponse api.GetAssociatedCollectors
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/ingestBudgets/" + id + "/collectors"
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

	err = json.Unmarshal(responseBody, &associatedCollectorsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	associatedCollectorResponseJson, err := json.MarshalIndent(associatedCollectorsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(associatedCollectorResponseJson))
	}
}
