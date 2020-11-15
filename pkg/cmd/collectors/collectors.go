package collectors

import (
	"github.com/spf13/cobra"
	cmdRoleList "github.com/wizedkyle/sumocli/pkg/cmd/roles/list"
)

func NewCmdCollector() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collectors <command>",
		Short: "Manage collectors",
		Long:  "Work with Sumo Logic collectors",
	}

	cmd.AddCommand(cmdRoleList.NewCmdRoleList())
	return cmd
}