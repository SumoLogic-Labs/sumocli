package delete

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
	response, err := client.DeleteIngestBudgetV2(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "The ingest budget was deleted successfully.")
	}
}
