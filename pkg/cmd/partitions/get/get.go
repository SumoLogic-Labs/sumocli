package get

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdPartitionsGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a partition with the given identifier from the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getPartition(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the partition")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getPartition(id string, client *cip.APIClient) {
	data, response, err := client.GetPartition(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
