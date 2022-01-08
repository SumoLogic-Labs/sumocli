package deactivate

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsDeactivate(client *cip.APIClient) *cobra.Command {
	var (
		organizationId     string
		parentDeploymentId string
	)
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deactivate an organization with the given identifier, deleting all its data and its subscription.",
		Run: func(cmd *cobra.Command, args []string) {
			deactivateOrganization(client, organizationId, parentDeploymentId)
		},
	}
	cmd.Flags().StringVar(&organizationId, "organizationId", "", "Specify the identifier of the organization for which the details are required.")
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.MarkFlagRequired("organizationId")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func deactivateOrganization(client *cip.APIClient, organizationId string, parentDeploymentId string) {
	_, response, err := client.DeleteOrganization(organizationId, parentDeploymentId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Organization was deactivated successfully.")
	}
}
