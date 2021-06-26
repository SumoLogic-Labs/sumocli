package delete

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdArchiveIngestionDelete() *cobra.Command {
	var (
		id       string
		sourceId int
	)

	cmd := &cobra.Command{
		Use: "delete",
		Short: "Delete an ingestion job with the given identifier from the organization. " +
			"The delete operation is only possible for jobs with a Succeeded or Failed status.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteArchiveIngestion(id, sourceId)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the archive source")
	cmd.Flags().IntVar(&sourceId, "sourceId", 0, "Specify the source Id of the Archive Source")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("sourceId")
	return cmd
}

func deleteArchiveIngestion(id string, sourceId int) {
	log := logging.GetConsoleLogger()
	sourceIdHex := fmt.Sprintf("%x", sourceId)
	requestUrl := "v1/archive/" + sourceIdHex + "/jobs/" + id
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
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
		fmt.Println("The ingestion job was deleted successfully.")
	}
}
