package create_organization_access_key

import (
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsCreateOrganizationAccessKey(client *cip.APIClient) *cobra.Command {
	var (
		corsHeaders        []string
		label              string
		organizationId     string
		parentDeploymentId string
	)
	cmd := &cobra.Command{
		Use:   "create-organization-access-key",
		Short: "Get an access ID and key pair for an existing organization based on the organization identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			createOrganizationAccessKey(client, corsHeaders, label, organizationId, parentDeploymentId)
		},
	}
	cmd.Flags().StringSliceVar(&corsHeaders, "corsHeaders", []string{}, "Specify a comma-separated list of domains for which the access key is valid.")
	cmd.Flags().StringVar(&label, "label", "", "Specify a name for the access key to be created.")
	cmd.Flags().StringVar(&organizationId, "organizationId", "", "Specify the identifier of the organization for which the details are required.")
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.MarkFlagRequired("label")
	cmd.MarkFlagRequired("organizationId")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func createOrganizationAccessKey(client *cip.APIClient, corsHeaders []string, label string, organizationId string,
	parentDeploymentId string) {
	if cmdutils.ValidateDeploymentId(parentDeploymentId) == false {
		fmt.Println(parentDeploymentId + " is not a valid deployment, https://help.sumologic.com/APIs/General-API-Information/Sumo-Logic-Endpoints-and-Firewall-Security provides a list of valid deployments.")
	}
	data, response, err := client.CreateOrganizationAccessKey(types.AccessKeyCreateRequest{
		Label:       label,
		CorsHeaders: corsHeaders,
	},
		parentDeploymentId,
		organizationId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
