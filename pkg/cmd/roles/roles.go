package roles

import (
	"github.com/spf13/cobra"
	cmdRoleCreate "github.com/wizedkyle/sumocli/pkg/cmd/roles/create"
	cmdRoleGet "github.com/wizedkyle/sumocli/pkg/cmd/roles/get"
	cmdRoleList "github.com/wizedkyle/sumocli/pkg/cmd/roles/list"
)

func NewCmdRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roles <command>",
		Short: "Manage roles",
		Long:  "Work with Sumo Logic roles",
	}

	cmd.AddCommand(cmdRoleCreate.NewCmdRoleCreate())
	cmd.AddCommand(cmdRoleGet.NewCmdRoleGet())
	cmd.AddCommand(cmdRoleList.NewCmdRoleList())
	return cmd
}
