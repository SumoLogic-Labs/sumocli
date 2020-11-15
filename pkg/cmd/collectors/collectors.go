package collectors

import (
	"github.com/spf13/cobra"
	cmdCollectorsList "github.com/wizedkyle/sumocli/pkg/cmd/collectors/list"
)

func NewCmdCollector() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collectors <command>",
		Short: "Manage collectors",
		Long:  "Work with Sumo Logic collectors",
	}

	cmd.AddCommand(cmdCollectorsList.NewCmdControllersList())
	return cmd
}