package assign_collector

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
	apiResponse, httpResponse, errorResponse := client.AssignCollectorToBudget(id, collectorId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
