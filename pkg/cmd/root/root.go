package root

import (
	"github.com/spf13/cobra"
	accessKeysCmd "github.com/wizedkyle/sumocli/pkg/cmd/access-keys"
	accountCmd "github.com/wizedkyle/sumocli/pkg/cmd/account"
	appsCmd "github.com/wizedkyle/sumocli/pkg/cmd/apps"
	azureCmd "github.com/wizedkyle/sumocli/pkg/cmd/azure"
	collectorCmd "github.com/wizedkyle/sumocli/pkg/cmd/collectors"
	contentCmd "github.com/wizedkyle/sumocli/pkg/cmd/content"
	dashboardsCmd "github.com/wizedkyle/sumocli/pkg/cmd/dashboards"
	foldersCmd "github.com/wizedkyle/sumocli/pkg/cmd/folders"
	ingestBudgetsCmd "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets"
	ingestBudgetsV2Cmd "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets-v2"
	liveTailCmd "github.com/wizedkyle/sumocli/pkg/cmd/live-tail"
	loginCmd "github.com/wizedkyle/sumocli/pkg/cmd/login"
	lookupTablesCmd "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables"
	passwordPolicyCmd "github.com/wizedkyle/sumocli/pkg/cmd/password-policy"
	permissionsCmd "github.com/wizedkyle/sumocli/pkg/cmd/permissions"
	roleCmd "github.com/wizedkyle/sumocli/pkg/cmd/roles"
	samlCmd "github.com/wizedkyle/sumocli/pkg/cmd/saml"
	serviceAllowlistCmd "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist"
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
	cmd.AddCommand(accountCmd.NewCmdAccount())
	cmd.AddCommand(accessKeysCmd.NewCmdAccessKeys())
	cmd.AddCommand(appsCmd.NewCmdApps())
	cmd.AddCommand(azureCmd.NewCmdAzure())
	cmd.AddCommand(collectorCmd.NewCmdCollectors())
	cmd.AddCommand(contentCmd.NewCmdContent())
	cmd.AddCommand(dashboardsCmd.NewCmdDashboards())
	cmd.AddCommand(foldersCmd.NewCmdFolders())
	cmd.AddCommand(ingestBudgetsCmd.NewCmdIngestBudgets())
	cmd.AddCommand(ingestBudgetsV2Cmd.NewCmdIngestBudgetsV2())
	cmd.AddCommand(liveTailCmd.NewCmdLiveTail())
	cmd.AddCommand(loginCmd.NewCmdLogin())
	cmd.AddCommand(lookupTablesCmd.NewCmdLookupTables())
	cmd.AddCommand(passwordPolicyCmd.NewCmdPasswordPolicy())
	cmd.AddCommand(permissionsCmd.NewCmdPermissions())
	cmd.AddCommand(roleCmd.NewCmdRole())
	cmd.AddCommand(samlCmd.NewCmdSaml())
	cmd.AddCommand(serviceAllowlistCmd.NewCmdServiceAllowlist())
	cmd.AddCommand(sourcesCmd.NewCmdSources())
	cmd.AddCommand(tokensCmd.NewCmdTokens())
	cmd.AddCommand(usersCmd.NewCmdUser())
	cmd.AddCommand(version.NewCmdVersion())

	// Add global, persistent flags - these apply for all commands and their subcommands

	return cmd
}
