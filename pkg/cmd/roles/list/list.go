package list

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdRoleList(client *cip.APIClient) *cobra.Command {
	var (
		limit  int32
		name   string
		sortBy bool
	)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic roles",
		Run: func(cmd *cobra.Command, args []string) {
			listRoles(client, limit, name, sortBy)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the role you want to retrieve.")
	cmd.Flags().BoolVar(&sortBy, "sortBy", false, "Sorts the roles by the name field.")
	return cmd
}

func listRoles(client *cip.APIClient, limit int32, name string, sortBy bool) {
	var options types.ListRolesOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	if sortBy == true {
		options.SortBy = optional.NewString("name")
	}
	if name != "" {
		options.Name = optional.NewString(name)
	}
	apiResponse, httpResponse, errorResponse := client.ListRoles(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listRolesPagination(client, options, paginationToken)
		paginationToken = apiResponse.Next
	}
}

func listRolesPagination(client *cip.APIClient, options types.ListRolesOpts, token string) types.ListRoleModelsResponse {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListRoles(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
