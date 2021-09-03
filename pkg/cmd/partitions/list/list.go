package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdPartitionsList(client *cip.APIClient) *cobra.Command {
	var (
		limit     int32
		viewTypes []string
	)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all partitions in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			listPartitions(limit, viewTypes, client)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	cmd.Flags().StringSliceVar(&viewTypes, "viewTypes", []string{}, "Specify the type of partitions to retrieve. "+
		"Valid values are: DefaultView, Partition, and AuditIndex. You can add multiple types for example DefaultView,Partition.")
	return cmd
}

func listPartitions(limit int32, viewTypes []string, client *cip.APIClient) {
	var options types.PartitionOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	options.ViewTypes = optional.NewInterface(viewTypes)
	apiResponse, httpResponse, errorResponse := client.ListPartitions(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	paginationToken = apiResponse.Next
	for paginationToken != "" {
		apiResponse = listPartitionsPagination(client, options, paginationToken)
		paginationToken = apiResponse.Next
	}
}

func listPartitionsPagination(client *cip.APIClient, options types.PartitionOpts, token string) types.ListPartitionsResponse {
	options.Token = optional.NewString(token)
	apiResponse, httpResponse, errorResponse := client.ListPartitions(&options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
	return apiResponse
}
