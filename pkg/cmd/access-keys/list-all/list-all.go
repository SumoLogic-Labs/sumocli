package list_all

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdAccessKeysListAll(client *cip.APIClient) *cobra.Command {
	var limit int32
	cmd := &cobra.Command{
		Use:   "list-all",
		Short: "List all access keys in your account.",
		Run: func(cmd *cobra.Command, args []string) {
			listAllAccessKeys(limit, client)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of access keys returned")
	return cmd
}

func listAllAccessKeys(limit int32, client *cip.APIClient) {
	var options types.AccessKeyOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	apiResponse, httpResponse, errorResponse := client.ListAccessKeys(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listAllAccessKeysPagination(client, options, paginationToken)
		paginationToken = apiResponse.Next
	}
}

func listAllAccessKeysPagination(client *cip.APIClient, options types.AccessKeyOpts, token string) types.PaginatedListAccessKeysResult {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListAccessKeys(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
