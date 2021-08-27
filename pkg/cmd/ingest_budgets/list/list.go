package list

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
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
	apiResponse, httpResponse, errorResponse := client.ListIngestBudgets(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listIngestBudgetsPagination(client, options, paginationToken)
		paginationToken = apiResponse.Next
	}
}

func listIngestBudgetsPagination(client *cip.APIClient, options types.ListIngestBudgetV1Opts, token string) types.ListIngestBudgetsResponse {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListIngestBudgets(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
