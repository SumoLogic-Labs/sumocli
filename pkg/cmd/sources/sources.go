package sources

import (
	"github.com/spf13/cobra"
	cmdSourcesCreate "github.com/wizedkyle/sumocli/pkg/cmd/sources/create"
	cmdSourcesList "github.com/wizedkyle/sumocli/pkg/cmd/sources/list"
)

func NewCmdSources() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sources <command>",
		Short: "Manages sources assigned to collectors",
	}

	cmd.AddCommand(cmdSourcesCreate.NewCmdCreateSource())
	cmd.AddCommand(cmdSourcesList.NewCmdSourceList())
	return cmd
}
