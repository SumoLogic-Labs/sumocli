package get_associated_collectors

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgetsGetAssociatedCollectors(client *cip.APIClient) *cobra.Command {
	var (
		id    string
		limit int32
	)
	cmd := &cobra.Command{
		Use:   "get-associated-collectors",
		Short: "Get a list of Collectors assigned to an ingest budget.",
		Run: func(cmd *cobra.Command, args []string) {
			getAssociatedCollectors(id, limit, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getAssociatedCollectors(id string, limit int32, client *cip.APIClient) {
	var options types.ListIngestBudgetV1Opts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	data, response, err := client.GetAssignedCollectors(id, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = getAssociatedCollectorsPagination(client, id, options, paginationToken)
		paginationToken = data.Next
	}
}

func getAssociatedCollectorsPagination(client *cip.APIClient, id string, options types.ListIngestBudgetV1Opts, token string) types.ListCollectorIdentitiesResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.GetAssignedCollectors(id, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
