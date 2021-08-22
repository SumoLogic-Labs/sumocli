package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
