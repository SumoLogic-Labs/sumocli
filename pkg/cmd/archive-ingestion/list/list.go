package list

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdArchiveIngestionList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all Archive Sources with the count and status of ingestion jobs.",
		Run: func(cmd *cobra.Command, args []string) {
			listArchiveIngestion()
		},
	}
	return cmd
}

func listArchiveIngestion() {
	var archiveIngestionResponse api.ListArchiveIngestion
	log := logging.GetConsoleLogger()
	requestUrl := "v1/archive/jobs/count"
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

	err = json.Unmarshal(responseBody, &archiveIngestionResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	archiveIngestionResponseJson, err := json.MarshalIndent(archiveIngestionResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(archiveIngestionResponseJson))
	}
}
