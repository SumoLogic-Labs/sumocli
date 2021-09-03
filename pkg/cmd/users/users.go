package users

import (
	cmdUserChange "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/users/change_email"
	cmdUserCreate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/users/create"
	cmdUserDelete "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/users/delete"
	cmdUserDisable "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/users/disable_mfa"
	cmdUserGet "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/users/get"
	cmdUserList "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/users/list"
	cmduserReset "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/users/reset_password"
	cmdUserUnlock "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/users/unlock"
	cmdUserUpdate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/users/update"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdUser(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users <command>",
		Short: "Manage users",
		Long:  "Work with Sumo Logic users",
	}

	cmd.AddCommand(cmdUserChange.NewCmdUserChangeEmail(client))
	cmd.AddCommand(cmdUserCreate.NewCmdUserCreate(client))
	cmd.AddCommand(cmdUserDelete.NewCmdUserDelete(client))
	cmd.AddCommand(cmdUserDisable.NewCmdUserDisableMFA(client))
	cmd.AddCommand(cmdUserGet.NewCmdGetUser(client))
	cmd.AddCommand(cmdUserList.NewCmdUserList(client))
	cmd.AddCommand(cmduserReset.NewCmdUserResetPassword(client))
	cmd.AddCommand(cmdUserUnlock.NewCmdUnlockUser(client))
	cmd.AddCommand(cmdUserUpdate.NewCmdUserUpdate(client))
	return cmd
}
