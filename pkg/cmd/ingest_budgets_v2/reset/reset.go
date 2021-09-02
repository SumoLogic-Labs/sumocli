package reset

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdIngestBudgetsV2Reset(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset ingest budget's current usage to 0 before the scheduled reset time.",
		Run: func(cmd *cobra.Command, args []string) {
			resetIngestBudgetV2(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("id")
	return cmd
}

func resetIngestBudgetV2(id string, client *cip.APIClient) {
	httpResponse, errorResponse := client.ResetUsageV2(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Ingest budget's usage was reset successfully.")
	}
}
