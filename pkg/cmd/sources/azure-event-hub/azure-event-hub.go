package azure_event_hub

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/internal/config"
	NewCmdAzureEventHubCreate "github.com/wizedkyle/sumocli/pkg/cmd/sources/azure-event-hub/create"
)

func NewCmdAzureEventHubSource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "azure-event-hub",
		Short: "Manage Azure Event Hub sources",
	}
	config.AddAzureFlags(cmd)
	cmd.AddCommand(NewCmdAzureEventHubCreate.NewCmdAzureEventHubSourceCreate())
	return cmd
}
