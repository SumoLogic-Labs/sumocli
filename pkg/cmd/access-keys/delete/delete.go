package delete

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccessKeysDelete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes the access key with the given accessId.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteAccessKey(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the access key to delete")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteAccessKey(id string, client *cip.APIClient) {
	httpResponse, errorResponse := client.DeleteAccessKey(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Access key was deleted successfully")
	}
}
