package cancel_retention_update

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdPartitionsCancelRetentionUpdate(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "cancel-retention-update",
		Short: "Cancel update to retention of a partition for which retention was updated previously using sumocli partitions update and the reduceRetentionPeriodImmediately parameter was set to false",
		Run: func(cmd *cobra.Command, args []string) {
			cancelRetentionUpdate(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the partition")
	cmd.MarkFlagRequired("id")
	return cmd
}

func cancelRetentionUpdate(id string, client *cip.APIClient) {
	response, err := client.CancelRetentionUpdate(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "The retention update was cancelled successfully.")
	}
}
