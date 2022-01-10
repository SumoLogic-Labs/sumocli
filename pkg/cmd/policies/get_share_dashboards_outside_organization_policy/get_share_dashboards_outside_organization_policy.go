package get_share_dashboards_outside_organization_policy

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdGetShareDashboardsOutsideOrganizationPolicy(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-share-dashboards-outside-organization-policy",
		Short: "Get the Share Dashboards Outside Organization policy.",
		Long: "Get the Share Dashboards Outside Organization policy. This policy allows users to share the dashboard with view only privileges outside of the organization (capability must be enabled from the Roles page). " +
			"Disabling this policy will disable all dashboards that have been shared outside of the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getShareDashboardsOutsideOrganizationPolicy(client)
		},
	}
	return cmd
}

func getShareDashboardsOutsideOrganizationPolicy(client *cip.APIClient) {
	data, response, err := client.GetShareDashboardsOutsideOrganizationPolicy()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
