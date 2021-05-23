package sources

import (
	"github.com/spf13/cobra"
	cmdHttpSources "github.com/wizedkyle/sumocli/pkg/cmd/sources/http"
	cmdSourcesList "github.com/wizedkyle/sumocli/pkg/cmd/sources/list"
)

func NewCmdSources() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sources",
		Short: "Manages sources assigned to collectors",
	}
	cmd.AddCommand(cmdHttpSources.NewCmdHttpSources())
	cmd.AddCommand(cmdSourcesList.NewCmdSourceList())
	return cmd
}
