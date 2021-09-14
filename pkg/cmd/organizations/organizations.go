package organizations

import (
	cmdOrganizationsCreateOrganization "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/create_organization"
	cmdOrganizationsCreateOrganizationAccessKey "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/create_organization_access_key"
	cmdOrganizationsDeactivate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/deactivate"
	cmdOrganizationsGetDeployments "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/get_deployments"
	cmdOrganizationsGetOrganization "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/get_organization"
	cmdOrganizationsGetOrganizationUsage "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/get_organization_usage"
	cmdOrganizationsGetParentOrganizationDetails "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/get_parent_organization_details"
	cmdOrganizationsGetParentOrganizationInfo "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/get_parent_organization_info"
	cmdOrganizationsGetSubdomainLoginUrl "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/get_subdomain_login_url"
	cmdOrganizationsList "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/list"
	cmdOrganizationsListOrganizationUsages "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/list_organization_usages"
	cmdOrganizationsUpdate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/organizations/update"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdOrganizations(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "organizations <command>",
		Short: "Manage organizations",
		Long:  "Commands that allow you to manage your Sumo Logic organizations.",
	}
	cmd.AddCommand(cmdOrganizationsCreateOrganization.NewCmdOrganizationsCreateOrganization(client))
	cmd.AddCommand(cmdOrganizationsCreateOrganizationAccessKey.NewCmdOrganizationsCreateOrganizationAccessKey(client))
	cmd.AddCommand(cmdOrganizationsDeactivate.NewCmdOrganizationsDeactivate(client))
	cmd.AddCommand(cmdOrganizationsGetDeployments.NewCmdOrganizationsGetDeployments(client))
	cmd.AddCommand(cmdOrganizationsGetOrganization.NewCmdOrganizationsGet(client))
	cmd.AddCommand(cmdOrganizationsGetOrganizationUsage.NewCmdOrganizationsGetOrganizationUsage(client))
	cmd.AddCommand(cmdOrganizationsGetParentOrganizationDetails.NewCmdOrganizationsGetParentOrganizationDetails(client))
	cmd.AddCommand(cmdOrganizationsGetParentOrganizationInfo.NewCmdOrganizationsGetParentOrganizationInfo(client))
	cmd.AddCommand(cmdOrganizationsGetSubdomainLoginUrl.NewCmdOrganizationsGetSubdomainLoginUrl(client))
	cmd.AddCommand(cmdOrganizationsList.NewCmdOrganizationsList(client))
	cmd.AddCommand(cmdOrganizationsListOrganizationUsages.NewCmdOrganizationsListOrganizationUsages(client))
	cmd.AddCommand(cmdOrganizationsUpdate.NewCmdOrganizationsUpdate(client))
	return cmd
}
