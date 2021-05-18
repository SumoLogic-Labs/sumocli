package reset

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdIngestBudgetsReset() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset ingest budget's current usage to 0 before the scheduled reset time.",
		Run: func(cmd *cobra.Command, args []string) {
			resetIngestBudget(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("id")
	return cmd
}

func resetIngestBudget(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/ingestBudgets/" + id + "/usage/reset"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		err := json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
	} else {
		fmt.Println("Ingest budget's usage was reset successfully.")
	}
}
