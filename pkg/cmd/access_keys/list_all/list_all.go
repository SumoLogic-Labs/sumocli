package list_all

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
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
	data, response, err := client.ListAccessKeys(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listAllAccessKeysPagination(client, options, paginationToken)
		paginationToken = data.Next
	}
}

func listAllAccessKeysPagination(client *cip.APIClient, options types.AccessKeyOpts, token string) types.PaginatedListAccessKeysResult {
	options.Token = optional.NewString(token)
	data, response, err := client.ListAccessKeys(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
