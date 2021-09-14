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
	data, response, err := client.ListDynamicParsingRules(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listDynamicRulesPagination(client, options, paginationToken)
		paginationToken = data.Next
	}
}

func listDynamicRulesPagination(client *cip.APIClient, options types.DynamicParsingRuleOpts, token string) types.ListDynamicRulesResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.ListDynamicParsingRules(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
