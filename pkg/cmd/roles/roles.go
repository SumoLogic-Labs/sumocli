package roles

import (
	"github.com/spf13/cobra"
	cmdRoleAssign "github.com/wizedkyle/sumocli/pkg/cmd/roles/assign"
	cmdRoleCreate "github.com/wizedkyle/sumocli/pkg/cmd/roles/create"
	cmdRoleDelete "github.com/wizedkyle/sumocli/pkg/cmd/roles/delete"
	cmdRoleGet "github.com/wizedkyle/sumocli/pkg/cmd/roles/get"
	cmdRoleList "github.com/wizedkyle/sumocli/pkg/cmd/roles/list"
	cmdRoleRemove "github.com/wizedkyle/sumocli/pkg/cmd/roles/remove"
	cmdRoleUpdate "github.com/wizedkyle/sumocli/pkg/cmd/roles/update"
)

func NewCmdRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roles <command>",
		Short: "Manage roles",
		Long:  "Work with Sumo Logic roles",
	}

	cmd.AddCommand(cmdRoleAssign.NewCmdRoleAssign())
	cmd.AddCommand(cmdRoleCreate.NewCmdRoleCreate())
	cmd.AddCommand(cmdRoleDelete.NewCmdRoleDelete())
	cmd.AddCommand(cmdRoleGet.NewCmdRoleGet())
	cmd.AddCommand(cmdRoleList.NewCmdRoleList())
	cmd.AddCommand(cmdRoleRemove.NewCmdRoleRemoveUser())
	cmd.AddCommand(cmdRoleUpdate.NewCmdRoleUpdate())
	return cmd
}
