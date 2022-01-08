package list_organization_usages

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsListOrganizationUsages(client *cip.APIClient) *cobra.Command {
	var (
		limit              int32
		organizations      []string
		parentDeploymentId string
	)
	cmd := &cobra.Command{
		Use:   "list-organization-usages",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			listOrganizationUsages(client, limit, organizations, parentDeploymentId)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results to return.")
	cmd.Flags().StringSliceVar(&organizations, "organizations", []string{}, "Specify a comma separated list of organizations you want usage details for.")
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func listOrganizationUsages(client *cip.APIClient, limit int32, organizations []string, parentDeploymentId string) {
	var (
		options         types.OrganizationsUsagesOpts
		paginationToken string
	)
	options.Limit = optional.NewInt32(limit)
	data, response, err := client.GetOrganizationUsages(organizations, parentDeploymentId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listOrganizationUsagesPagination(client, options, organizations, parentDeploymentId, paginationToken)
		paginationToken = data.Next
	}
}

func listOrganizationUsagesPagination(client *cip.APIClient, options types.OrganizationsUsagesOpts, organizations []string,
	parentDeploymentId string, token string) types.ListUsagesResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.GetOrganizationUsages(organizations, parentDeploymentId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
