package root

import (
	"github.com/spf13/cobra"
	azureCmd "github.com/wizedkyle/sumocli/pkg/cmd/azure"
	collectorCmd "github.com/wizedkyle/sumocli/pkg/cmd/collectors"
	contentCmd "github.com/wizedkyle/sumocli/pkg/cmd/content"
	foldersCmd "github.com/wizedkyle/sumocli/pkg/cmd/folders"
	liveTailCmd "github.com/wizedkyle/sumocli/pkg/cmd/live-tail"
	loginCmd "github.com/wizedkyle/sumocli/pkg/cmd/login"
	lookupTablesCmd "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables"
	roleCmd "github.com/wizedkyle/sumocli/pkg/cmd/roles"
	sourcesCmd "github.com/wizedkyle/sumocli/pkg/cmd/sources"
	tokensCmd "github.com/wizedkyle/sumocli/pkg/cmd/tokens"
	usersCmd "github.com/wizedkyle/sumocli/pkg/cmd/users"
	"github.com/wizedkyle/sumocli/pkg/cmd/version"
)

func NewCmdRoot() *cobra.Command {

	cmd := &cobra.Command{
		Use:              "sumocli <command> <subcommand> [flags]",
		Short:            "Sumo Logic CLI",
		Long:             "Interact with and manage Sumo Logic and Cloud SIEM Enterprise from the command line.",
		TraverseChildren: true,
	}

	// Add subcommands
	cmd.AddCommand(azureCmd.NewCmdAzure())
	cmd.AddCommand(collectorCmd.NewCmdCollectors())
	cmd.AddCommand(contentCmd.NewCmdContent())
	cmd.AddCommand(foldersCmd.NewCmdFolders())
	cmd.AddCommand(liveTailCmd.NewCmdLiveTail())
	cmd.AddCommand(loginCmd.NewCmdLogin())
	cmd.AddCommand(lookupTablesCmd.NewCmdLookupTables())
	cmd.AddCommand(roleCmd.NewCmdRole())
	cmd.AddCommand(sourcesCmd.NewCmdSources())
	cmd.AddCommand(tokensCmd.NewCmdTokens())
	cmd.AddCommand(usersCmd.NewCmdUser())
	cmd.AddCommand(version.NewCmdVersion())

	// Add global, persistent flags - these apply for all commands and their subcommands
	cmd.PersistentFlags().BoolP("verbose", "v", false, "Log with the highest level of verbosity available. Off by default.")
	cmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress any levelled log messages. General command output will still output to console. Off by default.")

	return cmd
}
