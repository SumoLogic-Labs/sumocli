package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdIngestBudgetsGet(client *cip.APIClient) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get an ingest budget by the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getIngestBudget(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getIngestBudget(id string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetIngestBudget(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
