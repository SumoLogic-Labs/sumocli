package get_parent_organization_info

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsGetParentOrganizationInfo(client *cip.APIClient) *cobra.Command {
	var parentDeploymentId string
	cmd := &cobra.Command{
		Use:   "get-parent-organization-info",
		Short: "Get information about parent organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getParentOrganizationInfo(client, parentDeploymentId)
		},
	}
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func getParentOrganizationInfo(client *cip.APIClient, parentDeploymentId string) {
	data, response, err := client.GetParentOrganizationInfo(parentDeploymentId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
