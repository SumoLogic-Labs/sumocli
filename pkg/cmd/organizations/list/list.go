package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsList(client *cip.APIClient) *cobra.Command {
	var (
		limit              int32
		parentDeploymentId string
		status             string
	)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all organizations in this account.",
		Run: func(cmd *cobra.Command, args []string) {
			list(client, limit, parentDeploymentId, status)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results to return.")
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.Flags().StringVar(&status, "status", "Active", "Specify the status of an organization, based on its subscription. Valid values are 'Active', 'Inactive', and 'All'. By default, only active organizations are listed.")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func list(client *cip.APIClient, limit int32, parentDeploymentId string, status string) {
	var (
		options         types.ListOrganizationsOpts
		paginationToken string
	)
	options.Limit = optional.NewInt32(limit)
	options.Status = optional.NewString(status)
	data, response, err := client.ListOrganizations(parentDeploymentId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listPagination(client, options, parentDeploymentId, paginationToken)
		paginationToken = data.Next
	}
}

func listPagination(client *cip.APIClient, options types.ListOrganizationsOpts, parentDeploymentId string, token string) types.ListOrganizationResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.ListOrganizations(parentDeploymentId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
