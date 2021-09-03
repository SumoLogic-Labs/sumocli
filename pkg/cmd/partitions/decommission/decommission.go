package decommission

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdPartitionsDecommission(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "decommission",
		Short: "Decommission a partition with the given identifier from the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			decommissionPartition(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the partition")
	cmd.MarkFlagRequired("id")
	return cmd
}

func decommissionPartition(id string, client *cip.APIClient) {
	httpResponse, errorResponse := client.DecommissionPartition(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "The partition was decommissioned successfully.")
	}
}
