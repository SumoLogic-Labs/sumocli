package remove_collector

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
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
	apiResponse, httpResponse, errorResponse := client.RemoveCollectorFromBudget(id, collectorId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
