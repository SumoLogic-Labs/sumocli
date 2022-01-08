package remove_collector

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmIngestBudgetsRemoveAssociatedCollector(client *cip.APIClient) *cobra.Command {
	var (
		collectorId string
		id          string
	)
	cmd := &cobra.Command{
		Use:   "remove-collector",
		Short: "Remove Collector from a budget.",
		Run: func(cmd *cobra.Command, args []string) {
			removeAssociatedCollector(collectorId, id, client)
		},
	}
	cmd.Flags().StringVar(&collectorId, "collectorId", "", "Specify the collector id to remove from the ingest budget")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("id")
	return cmd
}

func removeAssociatedCollector(collectorId string, id string, client *cip.APIClient) {
	data, response, err := client.RemoveCollectorFromBudget(id, collectorId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
