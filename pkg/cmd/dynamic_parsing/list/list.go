package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdDynamicParsingList(client *cip.APIClient) *cobra.Command {
	var limit int32
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all dynamic parsing rules.",
		Run: func(cmd *cobra.Command, args []string) {
			listDynamicParsingRules(limit, client)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	return cmd
}

func listDynamicParsingRules(limit int32, client *cip.APIClient) {
	var options types.DynamicParsingRuleOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	apiResponse, httpResponse, errorResponse := client.ListDynamicParsingRules(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listDynamicRulesPagination(client, options, paginationToken)
		paginationToken = apiResponse.Next
	}
}

func listDynamicRulesPagination(client *cip.APIClient, options types.DynamicParsingRuleOpts, token string) types.ListDynamicRulesResponse {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListDynamicParsingRules(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
