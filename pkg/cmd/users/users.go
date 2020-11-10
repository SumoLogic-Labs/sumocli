package users

import (
	"github.com/spf13/cobra"
	cmdUserCreate "github.com/wizedkyle/sumocli/pkg/cmd/users/create"
)

func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users <command>",
		Short: "Manage users",
		Long:  "Work with Sumo Logic users",
	}

	cmd.AddCommand(cmdUserCreate.NewCmdUserCreate())
	return cmd
}
