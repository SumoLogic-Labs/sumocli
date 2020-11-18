package users

import (
	"github.com/spf13/cobra"
	cmdUserCreate "github.com/wizedkyle/sumocli/pkg/cmd/users/create"
	cmdUserDelete "github.com/wizedkyle/sumocli/pkg/cmd/users/delete"
	cmdUserGet "github.com/wizedkyle/sumocli/pkg/cmd/users/get"
	cmdUserList "github.com/wizedkyle/sumocli/pkg/cmd/users/list"
	cmdUserUnlock "github.com/wizedkyle/sumocli/pkg/cmd/users/unlock"
)

func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users <command>",
		Short: "Manage users",
		Long:  "Work with Sumo Logic users",
	}

	cmd.AddCommand(cmdUserCreate.NewCmdUserCreate())
	cmd.AddCommand(cmdUserDelete.NewCmdUserDelete())
	cmd.AddCommand(cmdUserGet.NewCmdGetUser())
	cmd.AddCommand(cmdUserList.NewCmdUserList())
	cmd.AddCommand(cmdUserUnlock.NewCmdUnlockUser())
	return cmd
}
