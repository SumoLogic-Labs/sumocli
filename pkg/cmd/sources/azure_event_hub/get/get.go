package get

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAzureEventHubSourceGet(client *cip.APIClient) *cobra.Command {
	var (
		collectorId string
		sourceId    string
	)
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets information about an Azure Event Hub source",
		Run: func(cmd *cobra.Command, args []string) {
			getEventHubSource(collectorId, sourceId, client)
		},
	}
	cmd.Flags().StringVar(&collectorId, "collectorId", "", "Specify the collector id that the source is associated to")
	cmd.Flags().StringVar(&sourceId, "sourceId", "", "Specify the identifier of the source")
	return cmd
}

func getEventHubSource(collectorId string, sourceId string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetEventHubSource(collectorId, sourceId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
