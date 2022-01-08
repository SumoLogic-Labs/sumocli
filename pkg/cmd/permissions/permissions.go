package permissions

import (
	cmdPermissionsAdd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/permissions/add"
	cmdPermissionsGet "github.com/SumoLogic-Labs/sumocli/pkg/cmd/permissions/get"
	cmdPermissionsRemove "github.com/SumoLogic-Labs/sumocli/pkg/cmd/permissions/remove"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdPermissions(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "permissions",
		Short: "Manage permissions",
		Long:  "Commands that allow you to share your folders, searches, and dashboards with specific users or roles.",
	}
	cmd.AddCommand(cmdPermissionsAdd.NewCmdPermissionsAdd(client))
	cmd.AddCommand(cmdPermissionsGet.NewCmdPermissionsGet(client))
	cmd.AddCommand(cmdPermissionsRemove.NewCmdPermissionsRemove(client))
	return cmd
}
