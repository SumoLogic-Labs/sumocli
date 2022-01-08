package reset

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgetsReset(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset ingest budget's current usage to 0 before the scheduled reset time.",
		Run: func(cmd *cobra.Command, args []string) {
			resetIngestBudget(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("id")
	return cmd
}

func resetIngestBudget(id string, client *cip.APIClient) {
	response, err := client.ResetUsage(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Ingest budget's usage was reset successfully.")
	}
}
