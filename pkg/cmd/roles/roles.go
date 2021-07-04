package roles

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	cmdRoleAssign "github.com/wizedkyle/sumocli/pkg/cmd/roles/assign"
	cmdRoleCreate "github.com/wizedkyle/sumocli/pkg/cmd/roles/create"
	cmdRoleDelete "github.com/wizedkyle/sumocli/pkg/cmd/roles/delete"
	cmdRoleGet "github.com/wizedkyle/sumocli/pkg/cmd/roles/get"
	cmdRoleList "github.com/wizedkyle/sumocli/pkg/cmd/roles/list"
	cmdRoleRemove "github.com/wizedkyle/sumocli/pkg/cmd/roles/remove"
	cmdRoleUpdate "github.com/wizedkyle/sumocli/pkg/cmd/roles/update"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdRole(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roles <command>",
		Short: "Manage roles",
		Long:  "Commands that allow you to manage roles in your Sumo Logic tenant",
	}

	cmd.AddCommand(cmdRoleAssign.NewCmdRoleAssign(client, log))
	cmd.AddCommand(cmdRoleCreate.NewCmdRoleCreate(client, log))
	cmd.AddCommand(cmdRoleDelete.NewCmdRoleDelete(client, log))
	cmd.AddCommand(cmdRoleGet.NewCmdRoleGet(client, log))
	cmd.AddCommand(cmdRoleList.NewCmdRoleList(client, log))
	cmd.AddCommand(cmdRoleRemove.NewCmdRoleRemoveUser(client, log))
	cmd.AddCommand(cmdRoleUpdate.NewCmdRoleUpdate(client, log))
	return cmd
}
