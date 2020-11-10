package root

import (
	"github.com/spf13/cobra"
	loginCmd "github.com/wizedkyle/sumocli/pkg/cmd/login"
	roleCmd "github.com/wizedkyle/sumocli/pkg/cmd/roles"
	usersCmd "github.com/wizedkyle/sumocli/pkg/cmd/users"
	"github.com/wizedkyle/sumocli/pkg/cmd/version"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sumocli <command> <subcommand> [flags]",
		Short: "Sumo Logic CLI",
		Long:  "Manage Sumo Logic from the command line.",
	}

	// Add subcommands
	cmd.AddCommand(loginCmd.NewCmdLogin())
	cmd.AddCommand(roleCmd.NewCmdRole())
	cmd.AddCommand(usersCmd.NewCmdUser())
	cmd.AddCommand(version.NewCmdVersion())

	return cmd
}
