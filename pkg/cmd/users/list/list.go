package list

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
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
	var options types.UserManagementApiListUsersOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	if email != "" {
		options.Email = optional.NewString(email)
	}
	if sortBy != "" {
		options.SortBy = optional.NewString(sortBy)
	}
	apiResponse, httpResponse, errorResponse := client.ListUsers(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listUsersPagination(client, options, paginationToken)
		paginationToken = apiResponse.Next
	}
}

func listUsersPagination(client *cip.APIClient, options types.UserManagementApiListUsersOpts, token string) types.ListUserModelsResponse {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListUsers(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
