package root

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/config"
	accessKeysCmd "github.com/wizedkyle/sumocli/pkg/cmd/access-keys"
	accountCmd "github.com/wizedkyle/sumocli/pkg/cmd/account"
	appsCmd "github.com/wizedkyle/sumocli/pkg/cmd/apps"
	archiveIngestion "github.com/wizedkyle/sumocli/pkg/cmd/archive-ingestion"
	collectorCmd "github.com/wizedkyle/sumocli/pkg/cmd/collectors"
	contentCmd "github.com/wizedkyle/sumocli/pkg/cmd/content"
	dashboardsCmd "github.com/wizedkyle/sumocli/pkg/cmd/dashboards"
	dynamicParsingCmd "github.com/wizedkyle/sumocli/pkg/cmd/dynamic_parsing"
	fieldExtractionRulesCmd "github.com/wizedkyle/sumocli/pkg/cmd/field_extraction_rules"
	fieldManagement "github.com/wizedkyle/sumocli/pkg/cmd/field_management"
	foldersCmd "github.com/wizedkyle/sumocli/pkg/cmd/folders"
	healthEventsCmd "github.com/wizedkyle/sumocli/pkg/cmd/health_events"
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
)

func NewCmdRoot() *cobra.Command {

	cmd := &cobra.Command{
		Use:              "sumocli <command> <subcommand> [flags]",
		Short:            "Sumo Logic CLI",
		Long:             "Interact with and manage Sumo Logic and Cloud SIEM Enterprise from the command line.",
		TraverseChildren: true,
	}
	client := config.GetSumoLogicSDKConfig()
	// Add subcommands
	cmd.AddCommand(accountCmd.NewCmdAccount(client))
	cmd.AddCommand(accessKeysCmd.NewCmdAccessKeys(client))
	cmd.AddCommand(appsCmd.NewCmdApps(client))
	cmd.AddCommand(archiveIngestion.NewCmdArchiveIngestion(client))
	cmd.AddCommand(collectorCmd.NewCmdCollectors(client))
	cmd.AddCommand(contentCmd.NewCmdContent(client))
	cmd.AddCommand(dashboardsCmd.NewCmdDashboards(client))
	cmd.AddCommand(dynamicParsingCmd.NewCmdDynamicParsing(client))
	cmd.AddCommand(fieldExtractionRulesCmd.NewCmdFieldExtractionRules(client))
	cmd.AddCommand(fieldManagement.NewCmdFieldManagement(client))
	cmd.AddCommand(foldersCmd.NewCmdFolders(client))
	cmd.AddCommand(healthEventsCmd.NewCmdHealthEvents(client))
	cmd.AddCommand(ingestBudgetsCmd.NewCmdIngestBudgets())
	cmd.AddCommand(ingestBudgetsV2Cmd.NewCmdIngestBudgetsV2())
	cmd.AddCommand(liveTailCmd.NewCmdLiveTail())
	cmd.AddCommand(loginCmd.NewCmdLogin())
	cmd.AddCommand(lookupTablesCmd.NewCmdLookupTables())
	cmd.AddCommand(monitorsCmd.NewCmdMonitors())
	cmd.AddCommand(partitionsCmd.NewCmdPartitions())
	cmd.AddCommand(passwordPolicyCmd.NewCmdPasswordPolicy())
	cmd.AddCommand(permissionsCmd.NewCmdPermissions())
	cmd.AddCommand(roleCmd.NewCmdRole(client))
	cmd.AddCommand(samlCmd.NewCmdSaml())
	cmd.AddCommand(scheduledViewsCmd.NewCmdScheduledViews())
	cmd.AddCommand(serviceAllowlistCmd.NewCmdServiceAllowlist())
	cmd.AddCommand(sourcesCmd.NewCmdSources(client))
	cmd.AddCommand(tokensCmd.NewCmdTokens(client))
	cmd.AddCommand(usersCmd.NewCmdUser(client))
	cmd.AddCommand(version.NewCmdVersion())

	// Add global, persistent flags - these apply for all commands and their subcommands
	return cmd
}
