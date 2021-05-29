package sources

import (
	"github.com/spf13/cobra"
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
	cmd.AddCommand(cmdSourcesDelete.NewCmdDeleteSource())
	cmd.AddCommand(cmdHttpSources.NewCmdHttpSources())
	cmd.AddCommand(cmdSourcesList.NewCmdSourceList())
	cmd.AddCommand(cmdLocalFileSources.NewCmdLocalFileSources())
	return cmd
}
