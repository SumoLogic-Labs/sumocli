package get_parent_organization_details

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsGetParentOrganizationDetails(client *cip.APIClient) *cobra.Command {
	var parentDeploymentId string
	cmd := &cobra.Command{
		Use:   "get-parent-organization-details",
		Short: "Get the credits usage and allocation details of the current organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getParentOrganizationDetails(client, parentDeploymentId)
		},
	}
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func getParentOrganizationDetails(client *cip.APIClient, parentDeploymentId string) {
	data, response, err := client.GetParentOrganizationDetails(parentDeploymentId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
