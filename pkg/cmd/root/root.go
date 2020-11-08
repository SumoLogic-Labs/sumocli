package root

import (
	"github.com/spf13/cobra"
	loginCmd "github.com/wizedkyle/sumocli/pkg/cmd/login"
	roleCmd "github.com/wizedkyle/sumocli/pkg/cmd/roles"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sumocli <command> <subcommand> [flags]",
		Short: "Sumo Logic CLI",
		Long:  "Manage Sumo Logic from the command line.",
		//Example: heredoc.Doc(),
	}

	cmd.AddCommand(loginCmd.NewCmdRoleList())
	cmd.AddCommand(roleCmd.NewCmdRole())

	return cmd
}
