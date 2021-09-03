package created

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdPartitionCreate(client *cip.APIClient) *cobra.Command {
	var (
		name              string
		routingExpression string
		analyticsTier     string
		retentionPeriod   int32
		isCompliant       bool
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new partition.",
		Run: func(cmd *cobra.Command, args []string) {
			createPartition(name, routingExpression, analyticsTier, retentionPeriod, isCompliant, client)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the partition")
	cmd.Flags().StringVar(&routingExpression, "routingExpression", "", "Specify the query that defines the data to be included in the partition")
	cmd.Flags().StringVar(&analyticsTier, "analyticsTier", "continuous", "Specify the Data Tier where the data in the partition will reside. "+
		"Possible values are continuous, frequent, infrequent.")
	cmd.Flags().Int32Var(&retentionPeriod, "retentionPeriod", -1, "Specify the number of days to retain data in the partition. "+
		"-1 specifies that the default value for the account is used.")
	cmd.Flags().BoolVar(&isCompliant, "isCompliant", false, "Set to true if the partition is compliant")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("routingExpression")
	return cmd
}

func createPartition(name string, routingExpression string, analyticsTier string, retentionPeriod int32, isCompliant bool,
	client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.CreatePartition(types.CreatePartitionDefinition{
		Name:              name,
		RoutingExpression: routingExpression,
		AnalyticsTier:     analyticsTier,
		RetentionPeriod:   retentionPeriod,
		IsCompliant:       isCompliant,
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
