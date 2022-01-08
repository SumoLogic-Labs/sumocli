package get_organization_usage

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsGetOrganizationUsage(client *cip.APIClient) *cobra.Command {
	var (
		organizationId     string
		parentDeploymentId string
	)
	cmd := &cobra.Command{
		Use:   "get-organization-usage",
		Short: "Get the credits usage and allocation details of the current organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getOrganizationUsage(client, organizationId, parentDeploymentId)
		},
	}
	cmd.Flags().StringVar(&organizationId, "organizationId", "", "Specify the identifier of the organization for which the details are required.")
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.MarkFlagRequired("organizationId")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func getOrganizationUsage(client *cip.APIClient, organizationId string, parentDeploymentId string) {
	data, response, err := client.GetOrganizationUsage(organizationId, parentDeploymentId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
