package list_all

import (
	"github.com/antihax/optional"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdAccessKeysListAll(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var limit int32
	cmd := &cobra.Command{
		Use:   "list-all",
		Short: "List all access keys in your account.",
		Run: func(cmd *cobra.Command, args []string) {
			listAllAccessKeys(limit, client, log)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of access keys returned")
	return cmd
}

func listAllAccessKeys(limit int32, client *cip.APIClient, log *zerolog.Logger) {
	var options types.AccessKeyManagementApiListAccessKeysOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	apiResponse, httpResponse, errorResponse := client.ListAccessKeys(&options)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to list access keys")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listAllAccessKeysPagination(client, options, paginationToken, log)
		paginationToken = apiResponse.Next
	}
}

func listAllAccessKeysPagination(client *cip.APIClient, options types.AccessKeyManagementApiListAccessKeysOpts, token string, log *zerolog.Logger) types.PaginatedListAccessKeysResult {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListAccessKeys(&options)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to list access keys")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
