package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgetsList(client *cip.APIClient) *cobra.Command {
	var limit int32
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all ingest budgets.",
		Run: func(cmd *cobra.Command, args []string) {
			listIngestBudgets(limit, client)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	return cmd
}

func listIngestBudgets(limit int32, client *cip.APIClient) {
	var options types.ListIngestBudgetV1Opts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	data, response, err := client.ListIngestBudgets(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listIngestBudgetsPagination(client, options, paginationToken)
		paginationToken = data.Next
	}
}

func listIngestBudgetsPagination(client *cip.APIClient, options types.ListIngestBudgetV1Opts, token string) types.ListIngestBudgetsResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.ListIngestBudgets(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
