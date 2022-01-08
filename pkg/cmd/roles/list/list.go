package list

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
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
	var (
		options         types.ListRolesOpts
		paginationToken string
	)
	options.Limit = optional.NewInt32(limit)
	if sortBy == true {
		options.SortBy = optional.NewString("name")
	}
	if name != "" {
		options.Name = optional.NewString(name)
	}
	data, response, err := client.ListRoles(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listRolesPagination(client, options, paginationToken)
		paginationToken = data.Next
	}
}

func listRolesPagination(client *cip.APIClient, options types.ListRolesOpts, token string) types.ListRoleModelsResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.ListRoles(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
