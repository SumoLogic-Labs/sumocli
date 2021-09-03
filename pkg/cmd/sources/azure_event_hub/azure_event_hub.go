package azure_event_hub

import (
	NewCmdAzureEventHubCreate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/azure_event_hub/create"
	NewCmdAzureEventHubGet "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/azure_event_hub/get"
	NewCmdAzureEventHubUpdate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/azure_event_hub/update"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAzureEventHubSource(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use: "azure-event-hub",
		Short: "The Azure Event Hubs Source provides a secure endpoint to receive data from Azure Event Hubs. " +
			"It securely stores the required authentication, scheduling, and state tracking information. " +
			"This source is used to collect activity and resource logs from Azure.",
	}
	cmd.AddCommand(NewCmdAzureEventHubCreate.NewCmdAzureEventHubSourceCreate(client))
	cmd.AddCommand(NewCmdAzureEventHubGet.NewCmdAzureEventHubSourceGet(client))
	cmd.AddCommand(NewCmdAzureEventHubUpdate.NewCmdAzureEventHubSourceUpdate(client))
	return cmd
}
