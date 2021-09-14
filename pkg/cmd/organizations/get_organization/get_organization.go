package get_organization

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsGet(client *cip.APIClient) *cobra.Command {
	var (
		organizationId     string
		parentDeploymentId string
	)
	cmd := &cobra.Command{
		Use:   "get-organization",
		Short: "Get details of an existing organization based on an organization identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getOrganization(client, organizationId, parentDeploymentId)
		},
	}
	cmd.Flags().StringVar(&organizationId, "organizationId", "", "Specify the identifier of the organization for which the details are required.")
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.MarkFlagRequired("organizationId")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func getOrganization(client *cip.APIClient, organizationId string, parentDeploymentId string) {
	data, response, err := client.GetOrganization(organizationId, parentDeploymentId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
