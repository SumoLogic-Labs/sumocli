package azure_event_hub

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	NewCmdAzureEventHubCreate "github.com/wizedkyle/sumocli/pkg/cmd/sources/azure_event_hub/create"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAzureEventHubSource(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "azure-event-hub",
		Short: "Manage Azure Event Hub sources",
	}
	cmd.AddCommand(NewCmdAzureEventHubCreate.NewCmdAzureEventHubSourceCreate(client, log))
	return cmd
}
