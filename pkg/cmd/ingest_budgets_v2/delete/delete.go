package delete

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdIngestBudgetsV2Delete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an ingest budget with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteIngestBudgetV2(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteIngestBudgetV2(id string, client *cip.APIClient) {
	httpResponse, errorResponse := client.DeleteIngestBudgetV2(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "The ingest budget was deleted successfully.")
	}
}
