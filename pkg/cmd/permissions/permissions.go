package permissions

import (
	"github.com/spf13/cobra"
	cmdPermissionsAdd "github.com/wizedkyle/sumocli/pkg/cmd/permissions/add"
	cmdPermissionsGet "github.com/wizedkyle/sumocli/pkg/cmd/permissions/get"
	cmdPermissionsRemove "github.com/wizedkyle/sumocli/pkg/cmd/permissions/remove"
)

func NewCmdPermissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "permissions",
		Short: "Manage permissions",
		Long:  "Commands that allow you to share your folders, searches, and dashboards with specific users or roles.",
	}
	cmd.AddCommand(cmdPermissionsAdd.NewCmdPermissionsAdd())
	cmd.AddCommand(cmdPermissionsGet.NewCmdPermissionsGet())
	cmd.AddCommand(cmdPermissionsRemove.NewCmdPermissionsRemove())
	return cmd
}
