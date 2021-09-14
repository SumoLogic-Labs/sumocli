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
	data, response, err := client.ListPartitions(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listPartitionsPagination(client, options, paginationToken)
		paginationToken = data.Next
	}
}

func listPartitionsPagination(client *cip.APIClient, options types.PartitionOpts, token string) types.ListPartitionsResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.ListPartitions(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
