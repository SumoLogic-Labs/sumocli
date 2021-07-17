package create

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
	"time"
)

func NewCmdArchiveIngestionCreate(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var (
		endTime   string
		name      string
		sourceId  string
		startTime string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an ingestion job to pull data from your S3 bucket.",
		Run: func(cmd *cobra.Command, args []string) {
			createArchiveIngestion(endTime, name, sourceId, startTime, client, log)
		},
	}
	cmd.Flags().StringVar(&endTime, "endTime", "", "Specify the ending timestamp of the ingestion job. "+
		"The time format should be RFC3339 for example: 2021-06-01T12:00:00Z")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the ingestion job")
	cmd.Flags().StringVar(&sourceId, "sourceId", "", "Specify the source Id of the Archive Source for which "+
		"the job is to be added. You can create an S3 Archive source by running sumocli sources aws-s3-archive create.")
	cmd.Flags().StringVar(&startTime, "startTime", "", "Specify the starting timestamp of the ingestion job. "+
		"The time format should be RFC3339 for example: 2006-01-02T15:04:05Z07:00")
	cmd.MarkFlagRequired("endTime")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("sourceId")
	cmd.MarkFlagRequired("startTime")
	return cmd
}

func createArchiveIngestion(endTime string, name string, sourceId string, startTime string, client *cip.APIClient, log *zerolog.Logger) {
	format := "2006-01-02T15:04:05Z07:00"
	endTimeParsed, err := time.Parse(format, endTime)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse endTime, ensure time format is RFC3339 compliant")
	}
	startTimeParsed, err := time.Parse(format, startTime)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse startTime, ensure time format is RFC3339 compliant")
	}
	apiResponse, httpResponse, errorResponse := client.CreateArchiveJob(types.CreateArchiveJobRequest{
		Name:      name,
		StartTime: startTimeParsed,
		EndTime:   endTimeParsed,
	},
		sourceId)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to create archive job")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
