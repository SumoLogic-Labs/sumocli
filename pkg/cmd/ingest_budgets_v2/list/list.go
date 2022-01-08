package list

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgetsV2List(client *cip.APIClient) *cobra.Command {
	var limit int32
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all ingest budgets.",
		Run: func(cmd *cobra.Command, args []string) {
			listIngestBudgetsV2(limit, client)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	return cmd
}

func listIngestBudgetsV2(limit int32, client *cip.APIClient) {
	var options types.ListIngestBudgetV2Opts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	data, response, err := client.ListIngestBudgetsV2(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listIngestBudgetsV2Pagination(client, options, paginationToken)
		paginationToken = data.Next
	}
}

func listIngestBudgetsV2Pagination(client *cip.APIClient, options types.ListIngestBudgetV2Opts, token string) types.ListIngestBudgetsResponseV2 {
	options.Token = optional.NewString(token)
	data, response, err := client.ListIngestBudgetsV2(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
