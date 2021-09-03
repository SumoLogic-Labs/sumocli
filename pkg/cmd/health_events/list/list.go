package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdHealthEventsList(client *cip.APIClient) *cobra.Command {
	var limit int32
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all the unresolved health events in your account.",
		Run: func(cmd *cobra.Command, args []string) {
			listHealthEvents(limit, client)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of health events to return")
	return cmd
}

func listHealthEvents(limit int32, client *cip.APIClient) {
	var options types.HealthEventsOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	apiResponse, httpResponse, errorResponse := client.ListAllHealthEvents(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listHealthEventsPagination(client, options, paginationToken)
		paginationToken = apiResponse.Next
	}
}

func listHealthEventsPagination(client *cip.APIClient, options types.HealthEventsOpts, token string) types.ListHealthEventResponse {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListAllHealthEvents(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
