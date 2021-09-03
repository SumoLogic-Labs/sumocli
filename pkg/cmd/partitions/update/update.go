package update

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdPartitionUpdate(client *cip.APIClient) *cobra.Command {
	var (
		id                               string
		retentionPeriod                  int32
		reduceRetentionPeriodImmediately bool
		isCompliant                      bool
		routingExpression                string
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing partition in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			updatePartition(id, retentionPeriod, reduceRetentionPeriodImmediately, isCompliant, routingExpression, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the partition")
	cmd.Flags().Int32Var(&retentionPeriod, "retentionPeriod", -1, "Specify the number of days to retain data in the partition. "+
		"-1 specifies that the default value for the account is used.")
	cmd.Flags().BoolVar(&reduceRetentionPeriodImmediately, "reduceRetentionPeriodImmediately", false, "This is required if the newly specified retentionPeriod is less than the existing retention period. "+
		"A value of true says that data between the existing retention period and the new retention period should be deleted immediately. "+
		"If false, such data will be deleted after seven days.")
	cmd.Flags().BoolVar(&isCompliant, "isCompliant", false, "Set to true if the partition is compliant")
	cmd.Flags().StringVar(&routingExpression, "routingExpression", "", "Specify the query that defines the data to be included in the partition")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updatePartition(id string, retentionPeriod int32, reduceRetentionPeriodImmediately bool, isCompliant bool,
	routingExpression string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.UpdatePartition(types.UpdatePartitionDefinition{
		RetentionPeriod:                  retentionPeriod,
		ReduceRetentionPeriodImmediately: reduceRetentionPeriodImmediately,
		IsCompliant:                      isCompliant,
		RoutingExpression:                routingExpression,
	},
		id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
