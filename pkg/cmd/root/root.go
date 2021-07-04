package root

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/config"
	accessKeysCmd "github.com/wizedkyle/sumocli/pkg/cmd/access-keys"
	accountCmd "github.com/wizedkyle/sumocli/pkg/cmd/account"
	appsCmd "github.com/wizedkyle/sumocli/pkg/cmd/apps"
	archiveIngestion "github.com/wizedkyle/sumocli/pkg/cmd/archive-ingestion"
	azureCmd "github.com/wizedkyle/sumocli/pkg/cmd/azure"
	collectorCmd "github.com/wizedkyle/sumocli/pkg/cmd/collectors"
	contentCmd "github.com/wizedkyle/sumocli/pkg/cmd/content"
	dashboardsCmd "github.com/wizedkyle/sumocli/pkg/cmd/dashboards"
	dynamicParsingCmd "github.com/wizedkyle/sumocli/pkg/cmd/dynamic-parsing"
	fieldExtractionRulesCmd "github.com/wizedkyle/sumocli/pkg/cmd/field-extraction-rules"
	fieldManagement "github.com/wizedkyle/sumocli/pkg/cmd/field-management"
	foldersCmd "github.com/wizedkyle/sumocli/pkg/cmd/folders"
	healthEventsCmd "github.com/wizedkyle/sumocli/pkg/cmd/health-events"
	ingestBudgetsCmd "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets"
	ingestBudgetsV2Cmd "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets-v2"
	liveTailCmd "github.com/wizedkyle/sumocli/pkg/cmd/live-tail"
	loginCmd "github.com/wizedkyle/sumocli/pkg/cmd/login"
	lookupTablesCmd "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables"
	monitorsCmd "github.com/wizedkyle/sumocli/pkg/cmd/monitors"
	partitionsCmd "github.com/wizedkyle/sumocli/pkg/cmd/partitions"
	passwordPolicyCmd "github.com/wizedkyle/sumocli/pkg/cmd/password-policy"
	permissionsCmd "github.com/wizedkyle/sumocli/pkg/cmd/permissions"
	roleCmd "github.com/wizedkyle/sumocli/pkg/cmd/roles"
	samlCmd "github.com/wizedkyle/sumocli/pkg/cmd/saml"
	scheduledViewsCmd "github.com/wizedkyle/sumocli/pkg/cmd/scheduled-views"
	serviceAllowlistCmd "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist"
	sourcesCmd "github.com/wizedkyle/sumocli/pkg/cmd/sources"
	tokensCmd "github.com/wizedkyle/sumocli/pkg/cmd/tokens"
	usersCmd "github.com/wizedkyle/sumocli/pkg/cmd/users"
	"github.com/wizedkyle/sumocli/pkg/cmd/version"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdRoot() *cobra.Command {

	cmd := &cobra.Command{
		Use:              "sumocli <command> <subcommand> [flags]",
		Short:            "Sumo Logic CLI",
		Long:             "Interact with and manage Sumo Logic and Cloud SIEM Enterprise from the command line.",
		TraverseChildren: true,
	}
	client := config.GetSumoLogicSDKConfig()
	log := logging.GetLogger()
	// Add subcommands
	cmd.AddCommand(accountCmd.NewCmdAccount())
	cmd.AddCommand(accessKeysCmd.NewCmdAccessKeys())
	cmd.AddCommand(appsCmd.NewCmdApps())
	cmd.AddCommand(archiveIngestion.NewCmdArchiveIngestion())
	cmd.AddCommand(azureCmd.NewCmdAzure())
	cmd.AddCommand(collectorCmd.NewCmdCollectors())
	cmd.AddCommand(contentCmd.NewCmdContent())
	cmd.AddCommand(dashboardsCmd.NewCmdDashboards())
	cmd.AddCommand(dynamicParsingCmd.NewCmdDynamicParsing())
	cmd.AddCommand(fieldExtractionRulesCmd.NewCmdFieldExtractionRules())
	cmd.AddCommand(fieldManagement.NewCmdFieldManagement())
	cmd.AddCommand(foldersCmd.NewCmdFolders())
	cmd.AddCommand(healthEventsCmd.NewCmdHealthEvents())
	cmd.AddCommand(ingestBudgetsCmd.NewCmdIngestBudgets())
	cmd.AddCommand(ingestBudgetsV2Cmd.NewCmdIngestBudgetsV2())
	cmd.AddCommand(liveTailCmd.NewCmdLiveTail())
	cmd.AddCommand(loginCmd.NewCmdLogin())
	cmd.AddCommand(lookupTablesCmd.NewCmdLookupTables())
	cmd.AddCommand(monitorsCmd.NewCmdMonitors())
	cmd.AddCommand(partitionsCmd.NewCmdPartitions())
	cmd.AddCommand(passwordPolicyCmd.NewCmdPasswordPolicy())
	cmd.AddCommand(permissionsCmd.NewCmdPermissions())
	cmd.AddCommand(roleCmd.NewCmdRole(client, log))
	cmd.AddCommand(samlCmd.NewCmdSaml())
	cmd.AddCommand(scheduledViewsCmd.NewCmdScheduledViews())
	cmd.AddCommand(serviceAllowlistCmd.NewCmdServiceAllowlist())
	cmd.AddCommand(sourcesCmd.NewCmdSources())
	cmd.AddCommand(tokensCmd.NewCmdTokens())
	cmd.AddCommand(usersCmd.NewCmdUser())
	cmd.AddCommand(version.NewCmdVersion())

	// Add global, persistent flags - these apply for all commands and their subcommands
	return cmd
}
