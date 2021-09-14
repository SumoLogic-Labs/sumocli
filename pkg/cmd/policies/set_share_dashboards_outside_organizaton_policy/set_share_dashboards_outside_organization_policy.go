package set_share_dashboards_outside_organizaton_policy

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdSetShareDashboardsOutsideOrganizationPolicy(client *cip.APIClient) *cobra.Command {
	var enabled bool
	cmd := &cobra.Command{
		Use:   "set-share-dashboards-outside-organization-policy",
		Short: "Set the Share Dashboards Outside Organization policy.",
		Long: "Set the Share Dashboards Outside Organization policy. This policy allows users to share the dashboard with view only privileges outside of the organization (capability must be enabled from the Roles page). " +
			"Disabling this policy will disable all dashboards that have been shared outside of the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			setShareDashboardsOutsideOrganizationPolicy(client, enabled)
		},
	}
	cmd.Flags().BoolVar(&enabled, "enabled", true, "If set to false the audit policy is disabled.")
	return cmd
}

func setShareDashboardsOutsideOrganizationPolicy(client *cip.APIClient, enabled bool) {
	data, response, err := client.SetShareDashboardsOutsideOrganizationPolicy(types.ShareDashboardsOutsideOrganizationPolicy{
		enabled,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
