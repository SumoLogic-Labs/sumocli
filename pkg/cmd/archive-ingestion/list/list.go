package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdArchiveIngestionList(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all Archive Sources with the count and status of ingestion jobs.",
		Run: func(cmd *cobra.Command, args []string) {
			listArchiveIngestion(client)
		},
	}
	return cmd
}

func listArchiveIngestion(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.ListArchiveJobsCountPerSource()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
