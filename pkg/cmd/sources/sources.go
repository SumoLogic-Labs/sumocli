package sources

import (
	"github.com/spf13/cobra"
	cmdAzureEventHubSources "github.com/wizedkyle/sumocli/pkg/cmd/sources/azure-event-hub"
	cmdSourcesDelete "github.com/wizedkyle/sumocli/pkg/cmd/sources/delete"
	cmdHttpSources "github.com/wizedkyle/sumocli/pkg/cmd/sources/http"
	cmdSourcesList "github.com/wizedkyle/sumocli/pkg/cmd/sources/list"
	cmdLocalFileSources "github.com/wizedkyle/sumocli/pkg/cmd/sources/local-file"
)

func NewCmdSources() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sources",
		Short: "Manages sources assigned to collectors",
	}
	cmd.AddCommand(cmdAzureEventHubSources.NewCmdAzureEventHubSource())
	cmd.AddCommand(cmdSourcesDelete.NewCmdDeleteSource())
	cmd.AddCommand(cmdHttpSources.NewCmdHttpSources())
	cmd.AddCommand(cmdSourcesList.NewCmdSourceList())
	cmd.AddCommand(cmdLocalFileSources.NewCmdLocalFileSources())
	return cmd
}
