package list

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdTransformationRulesList(client *cip.APIClient) *cobra.Command {
	var limit int32
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of transformation rules in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			list(client, limit)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results, this is set to 100 by default.")
	return cmd
}

func list(client *cip.APIClient, limit int32) {
	var (
		options         types.TransformationRulesOpts
		paginationToken string
	)
	options.Limit = optional.NewInt32(limit)
	data, response, err := client.ListTransformationRules(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listPagination(client, options, paginationToken)
		paginationToken = data.Next
	}
}

func listPagination(client *cip.APIClient, options types.TransformationRulesOpts, token string) types.TransformationRulesResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.ListTransformationRules(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
