package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
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
	apiResponse, httpResponse, errorResponse := client.ListIngestBudgetsV2(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listIngestBudgetsV2Pagination(client, options, paginationToken)
		paginationToken = apiResponse.Next
	}
}

func listIngestBudgetsV2Pagination(client *cip.APIClient, options types.ListIngestBudgetV2Opts, token string) types.ListIngestBudgetsResponseV2 {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListIngestBudgetsV2(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
