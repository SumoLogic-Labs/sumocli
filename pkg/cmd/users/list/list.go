package list

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdUserList(client *cip.APIClient) *cobra.Command {
	var (
		email  string
		limit  int32
		sortBy string
	)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic users",
		Run: func(cmd *cobra.Command, args []string) {
			listUsers(email, limit, sortBy, client)
		},
	}
	cmd.Flags().StringVar(&email, "email", "", "Specify the email address of the user")
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results")
	cmd.Flags().StringVar(&sortBy, "sortBy", "", "Sort the results by firstName, lastName or email")
	return cmd
}

func listUsers(email string, limit int32, sortBy string, client *cip.APIClient) {
	var options types.ListUsersOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	if email != "" {
		options.Email = optional.NewString(email)
	}
	if sortBy != "" {
		options.SortBy = optional.NewString(sortBy)
	}
	data, response, err := client.ListUsers(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listUsersPagination(client, options, paginationToken)
		paginationToken = data.Next
	}
}

func listUsersPagination(client *cip.APIClient, options types.ListUsersOpts, token string) types.ListUserModelsResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.ListUsers(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
