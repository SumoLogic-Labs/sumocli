package reset

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
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
	httpResponse, errorResponse := client.ResetUsage(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Ingest budget's usage was reset successfully.")
	}
}
