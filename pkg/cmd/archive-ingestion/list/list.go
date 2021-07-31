package list

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
