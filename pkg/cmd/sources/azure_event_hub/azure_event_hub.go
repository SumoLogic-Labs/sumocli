package azure_event_hub

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	NewCmdAzureEventHubCreate "github.com/wizedkyle/sumocli/pkg/cmd/sources/azure_event_hub/create"
	NewCmdAzureEventHubGet "github.com/wizedkyle/sumocli/pkg/cmd/sources/azure_event_hub/get"
	NewCmdAzureEventHubUpdate "github.com/wizedkyle/sumocli/pkg/cmd/sources/azure_event_hub/update"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAzureEventHubSource(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use: "azure-event-hub",
		Short: "The Azure Event Hubs Source provides a secure endpoint to receive data from Azure Event Hubs. " +
			"It securely stores the required authentication, scheduling, and state tracking information. " +
			"This source is used to collect activity and resource logs from Azure.",
	}
	cmd.AddCommand(NewCmdAzureEventHubCreate.NewCmdAzureEventHubSourceCreate(client, log))
	cmd.AddCommand(NewCmdAzureEventHubGet.NewCmdAzureEventHubSourceGet(client, log))
	cmd.AddCommand(NewCmdAzureEventHubUpdate.NewCmdAzureEventHubSourceUpdate(client, log))
	return cmd
}
