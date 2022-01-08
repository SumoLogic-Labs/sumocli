package root

import (
	"github.com/SumoLogic-Labs/sumocli/config"
	accessKeysCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/access_keys"
	accountCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/account"
	appsCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/apps"
	archiveIngestion "github.com/SumoLogic-Labs/sumocli/pkg/cmd/archive-ingestion"
	collectorCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/collectors"
	ConfigureCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/configure"
	contentCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/content"
	dashboardsCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/dashboards"
	dynamicParsingCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/dynamic_parsing"
	fieldExtractionRulesCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_extraction_rules"
	fieldManagement "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management"
	foldersCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders"
	healthEventsCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/health_events"
	ingestBudgetsCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets"
	ingestBudgetsV2Cmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets_v2"
	liveTailCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/live-tail"
	lookupTablesCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables"
	monitorsCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/monitors"
	organizationsCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/organizations"
	partitionsCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/partitions"
	passwordPolicyCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/password_policy"
	permissionsCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/permissions"
	policiesCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies"
	roleCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/roles"
	samlCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/saml"
	scheduledViewsCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/scheduled-views"
	serviceAllowlistCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/service_allowlist"
	sourcesCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/sources"
	tokensCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/tokens"
	transformationRulesCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/transformation_rules"
	usersCmd "github.com/SumoLogic-Labs/sumocli/pkg/cmd/users"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmd/version"
	"github.com/spf13/cobra"
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
	cmd.AddCommand(ingestBudgetsCmd.NewCmdIngestBudgets(client))
	cmd.AddCommand(ingestBudgetsV2Cmd.NewCmdIngestBudgetsV2(client))
	cmd.AddCommand(liveTailCmd.NewCmdLiveTail())
	cmd.AddCommand(ConfigureCmd.NewCmdConfigure())
	cmd.AddCommand(lookupTablesCmd.NewCmdLookupTables(client))
	cmd.AddCommand(monitorsCmd.NewCmdMonitors())
	cmd.AddCommand(organizationsCmd.NewCmdOrganizations(client))
	cmd.AddCommand(partitionsCmd.NewCmdPartitions(client))
	cmd.AddCommand(passwordPolicyCmd.NewCmdPasswordPolicy(client))
	cmd.AddCommand(permissionsCmd.NewCmdPermissions(client))
	cmd.AddCommand(policiesCmd.NewCmdPolicies(client))
	cmd.AddCommand(roleCmd.NewCmdRole(client))
	cmd.AddCommand(samlCmd.NewCmdSaml())
	cmd.AddCommand(scheduledViewsCmd.NewCmdScheduledViews())
	cmd.AddCommand(serviceAllowlistCmd.NewCmdServiceAllowlist(client))
	cmd.AddCommand(sourcesCmd.NewCmdSources(client))
	cmd.AddCommand(tokensCmd.NewCmdTokens(client))
	cmd.AddCommand(transformationRulesCmd.NewCmdTransformationRules(client))
	cmd.AddCommand(usersCmd.NewCmdUser(client))
	cmd.AddCommand(version.NewCmdVersion())

	// Add global, persistent flags - these apply for all commands and their subcommands
	return cmd
}
