package reset

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
	response, err := client.ResetUsageV2(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Ingest budget's usage was reset successfully.")
	}
}
