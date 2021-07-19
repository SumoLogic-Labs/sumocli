package users

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	cmdUserChange "github.com/wizedkyle/sumocli/pkg/cmd/users/change_email"
	cmdUserCreate "github.com/wizedkyle/sumocli/pkg/cmd/users/create"
	cmdUserDelete "github.com/wizedkyle/sumocli/pkg/cmd/users/delete"
	cmdUserDisable "github.com/wizedkyle/sumocli/pkg/cmd/users/disable_mfa"
	cmdUserGet "github.com/wizedkyle/sumocli/pkg/cmd/users/get"
	cmdUserList "github.com/wizedkyle/sumocli/pkg/cmd/users/list"
	cmduserReset "github.com/wizedkyle/sumocli/pkg/cmd/users/reset_password"
	cmdUserUnlock "github.com/wizedkyle/sumocli/pkg/cmd/users/unlock"
	cmdUserUpdate "github.com/wizedkyle/sumocli/pkg/cmd/users/update"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdUser(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users <command>",
		Short: "Manage users",
		Long:  "Work with Sumo Logic users",
	}

	cmd.AddCommand(cmdUserChange.NewCmdUserChangeEmail(client, log))
	cmd.AddCommand(cmdUserCreate.NewCmdUserCreate(client, log))
	cmd.AddCommand(cmdUserDelete.NewCmdUserDelete(client, log))
	cmd.AddCommand(cmdUserDisable.NewCmdUserDisableMFA(client, log))
	cmd.AddCommand(cmdUserGet.NewCmdGetUser(client, log))
	cmd.AddCommand(cmdUserList.NewCmdUserList(client, log))
	cmd.AddCommand(cmduserReset.NewCmdUserResetPassword(client, log))
	cmd.AddCommand(cmdUserUnlock.NewCmdUnlockUser(client, log))
	cmd.AddCommand(cmdUserUpdate.NewCmdUserUpdate(client, log))
	return cmd
}
