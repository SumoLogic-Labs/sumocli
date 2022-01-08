package delete

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
	response, err := client.DeleteAccessKey(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Access key was deleted successfully")
	}
}
