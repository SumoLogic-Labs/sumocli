package roles

import (
	cmdRoleAssign "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/roles/assign"
	cmdRoleCreate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/roles/create"
	cmdRoleDelete "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/roles/delete"
	cmdRoleGet "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/roles/get"
	cmdRoleList "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/roles/list"
	cmdRoleRemove "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/roles/remove"
	cmdRoleUpdate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/roles/update"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdRole(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roles <command>",
		Short: "Manage roles",
		Long:  "Commands that allow you to manage roles in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdRoleAssign.NewCmdRoleAssign(client))
	cmd.AddCommand(cmdRoleCreate.NewCmdRoleCreate(client))
	cmd.AddCommand(cmdRoleDelete.NewCmdRoleDelete(client))
	cmd.AddCommand(cmdRoleGet.NewCmdRoleGet(client))
	cmd.AddCommand(cmdRoleList.NewCmdRoleList(client))
	cmd.AddCommand(cmdRoleRemove.NewCmdRoleRemoveUser(client))
	cmd.AddCommand(cmdRoleUpdate.NewCmdRoleUpdate(client))
	return cmd
}
