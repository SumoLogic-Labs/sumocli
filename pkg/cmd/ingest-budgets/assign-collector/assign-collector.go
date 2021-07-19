package assign_collector

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdIngestBudgetsAssignCollector() *cobra.Command {
	var (
		collectorId string
		id          string
	)

	cmd := &cobra.Command{
		Use:   "assign-collector",
		Short: "Assign a Collector to a budget.",
		Run: func(cmd *cobra.Command, args []string) {
			assignCollector(collectorId, id)
		},
	}
	cmd.Flags().StringVar(&collectorId, "collectorId", "", "Specify the collector id to add to the ingest budget")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("id")
	return cmd
}

func assignCollector(collectorId string, id string) {
	var ingestBudgetResponse api.GetIngestBudget
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/ingestBudgets/" + id + "/collectors/" + collectorId
	client, request := factory.NewHttpRequest("PUT", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &ingestBudgetResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	ingestBudgetResponseJson, err := json.MarshalIndent(ingestBudgetResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(ingestBudgetResponseJson))
	}
}
