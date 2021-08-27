package cancel_retention_update

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
	httpResponse, errorResponse := client.CancelRetentionUpdate(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "The retention update was cancelled successfully.")
	}
}
