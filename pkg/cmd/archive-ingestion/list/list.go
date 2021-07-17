package list

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdArchiveIngestionList(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all Archive Sources with the count and status of ingestion jobs.",
		Run: func(cmd *cobra.Command, args []string) {
			listArchiveIngestion(client, log)
		},
	}
	return cmd
}

func listArchiveIngestion(client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.ListArchiveJobsCountPerSource()
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to list archive jobs")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
