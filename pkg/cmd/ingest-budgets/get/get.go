package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdIngestBudgetsGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get an ingest budget by the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getIngestBudget(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getIngestBudget(id string) {
	var ingestBudgetResponse api.GetIngestBudget
	log := logging.GetConsoleLogger()
	requestUrl := "v1/ingestBudgets/" + id
	client, request := factory.NewHttpRequest("GET", requestUrl)
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
