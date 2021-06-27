package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdArchiveIngestionCreate() *cobra.Command {
	var (
		endTime   string
		name      string
		sourceId  int
		startTime string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an ingestion job to pull data from your S3 bucket.",
		Run: func(cmd *cobra.Command, args []string) {
			createArchiveIngestion(endTime, name, sourceId, startTime)
		},
	}
	cmd.Flags().StringVar(&endTime, "endTime", "", "Specify the ending timestamp of the ingestion job. "+
		"The time format should be RFC3339 for example: 2021-06-01T12:00:00Z")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the ingestion job")
	cmd.Flags().IntVar(&sourceId, "sourceId", 0, "Specify the source Id of the Archive Source for which "+
		"the job is to be added. You can create an S3 Archive source by running sumocli sources aws-s3-archive create.")
	cmd.Flags().StringVar(&startTime, "startTime", "", "Specify the starting timestamp of the ingestion job. "+
		"The time format should be RFC3339 for example: 2021-06-01T12:00:00Z")
	cmd.MarkFlagRequired("endTime")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("sourceId")
	cmd.MarkFlagRequired("startTime")
	return cmd
}

func createArchiveIngestion(endTime string, name string, sourceId int, startTime string) {
	var archiveIngestionResponse api.CreateArchiveIngestionResponse
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateArchiveIngestion{
		Name:      name,
		StartTime: startTime,
		EndTime:   endTime,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	sourceIdHex := fmt.Sprintf("%x", sourceId)
	requestUrl := "v1/archive/" + sourceIdHex + "/jobs"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
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
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(archiveIngestionResponseJson))
	}
}
