package assign_collector

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgetsAssignCollector(client *cip.APIClient) *cobra.Command {
	var (
		collectorId string
		id          string
	)
	cmd := &cobra.Command{
		Use:   "assign-collector",
		Short: "Assign a Collector to a budget.",
		Run: func(cmd *cobra.Command, args []string) {
			assignCollector(collectorId, id, client)
		},
	}
	cmd.Flags().StringVar(&collectorId, "collectorId", "", "Specify the collector id to add to the ingest budget")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("id")
	return cmd
}

func assignCollector(collectorId string, id string, client *cip.APIClient) {
	data, response, err := client.AssignCollectorToBudget(id, collectorId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
