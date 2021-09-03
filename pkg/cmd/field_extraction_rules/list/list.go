package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdFieldExtractionRulesList(client *cip.APIClient) *cobra.Command {
	var limit int32
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all field extraction rules.",
		Run: func(cmd *cobra.Command, args []string) {
			listFieldExtractionRules(limit, client)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	return cmd
}

func listFieldExtractionRules(limit int32, client *cip.APIClient) {
	var options types.ExtractionRuleOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	apiResponse, httpResponse, errorResponse := client.ListExtractionRules(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listFieldExtractionRulesPagination(client, options, paginationToken)
		paginationToken = apiResponse.Next
	}
}

func listFieldExtractionRulesPagination(client *cip.APIClient, options types.ExtractionRuleOpts, token string) types.ListExtractionRulesResponse {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListExtractionRules(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
