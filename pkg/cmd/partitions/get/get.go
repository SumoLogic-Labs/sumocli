package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
	apiResponse, httpResponse, errorResponse := client.GetPartition(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
