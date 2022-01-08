package policies

import (
	cmdPoliciesGetAuditPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/get_audit_policy"
	cmdPoliciesGetDataAccessLevelPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/get_data_access_level_policy"
	cmdPoliciesGetMaxUserSessionTimeoutPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/get_max_user_session_timeout_policy"
	cmdPoliciesGetSearchAuditPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/get_search_audit_policy"
	cmdPoliciesGetShareDashboardsOutsideOrganizationPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/get_share_dashboards_outside_organization_policy"
	cmdPoliciesGetUserConcurrentSessionsLimitPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/get_user_concurrent_sessions_limit_policy"
	cmdPoliciesSetAuditPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/set_audit_policy"
	cmdPoliciesSetDataAccessLevelPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/set_data_access_level_policy"
	cmdPoliciesSetMaxUserSessionTimeoutPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/set_max_user_session_timeout_policy"
	cmdPoliciesSetSearchAuditPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/set_search_audit_policy"
	cmdPoliciesSetShareDashboardsOutsideOrganizationPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/set_share_dashboards_outside_organizaton_policy"
	cmdPoliciesSetUserConcurrentSessionsLimitPolicy "github.com/SumoLogic-Labs/sumocli/pkg/cmd/policies/set_user_concurrent_sessions_limit_policy"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdPolicies(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policies <command>",
		Short: "Manage policies",
		Long:  "Commands that allow you to control the security and share settings of your organisation.",
	}
	cmd.AddCommand(cmdPoliciesGetAuditPolicy.NewCmdGetAuditPolicy(client))
	cmd.AddCommand(cmdPoliciesGetDataAccessLevelPolicy.NewCmdGetDataAccessLevelPolicy(client))
	cmd.AddCommand(cmdPoliciesGetMaxUserSessionTimeoutPolicy.NewCmdGetMaxUserSessionTimeoutPolicy(client))
	cmd.AddCommand(cmdPoliciesGetSearchAuditPolicy.NewCmdGetSearchAuditPolicy(client))
	cmd.AddCommand(cmdPoliciesGetShareDashboardsOutsideOrganizationPolicy.NewCmdGetShareDashboardsOutsideOrganizationPolicy(client))
	cmd.AddCommand(cmdPoliciesGetUserConcurrentSessionsLimitPolicy.NewCmdGetUserConcurrentSessionsLimitPolicy(client))
	cmd.AddCommand(cmdPoliciesSetAuditPolicy.NewCmdSetAuditPolicy(client))
	cmd.AddCommand(cmdPoliciesSetDataAccessLevelPolicy.NewCmdSetDataAccessLevelPolicy(client))
	cmd.AddCommand(cmdPoliciesSetMaxUserSessionTimeoutPolicy.NewCmdSetMaxUserSessionTimeoutPolicy(client))
	cmd.AddCommand(cmdPoliciesSetSearchAuditPolicy.NewCmdSetSearchAuditPolicy(client))
	cmd.AddCommand(cmdPoliciesSetShareDashboardsOutsideOrganizationPolicy.NewCmdSetShareDashboardsOutsideOrganizationPolicy(client))
	cmd.AddCommand(cmdPoliciesSetUserConcurrentSessionsLimitPolicy.NewCmdSetUserConcurrentSessionsLimitPolicy(client))
	return cmd
}
