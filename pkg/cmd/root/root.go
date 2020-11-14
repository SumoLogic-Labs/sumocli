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
		TraverseChildren: true,
	}

	// Add subcommands
	cmd.AddCommand(loginCmd.NewCmdLogin())
	cmd.AddCommand(roleCmd.NewCmdRole())
	cmd.AddCommand(usersCmd.NewCmdUser())
	cmd.AddCommand(version.NewCmdVersion())

	// Add global, persistent flags - these apply for all commands and their subcommands
	cmd.PersistentFlags().BoolP("verbose", "v", false, "Log with the highest level of verbosity available.")

	return cmd
}