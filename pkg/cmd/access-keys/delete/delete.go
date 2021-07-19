package delete

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccessKeysDelete(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes the access key with the given accessId.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteAccessKey(id, client, log)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the access key to delete")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteAccessKey(id string, client *cip.APIClient, log *zerolog.Logger) {
	httpResponse, errorResponse := client.DeleteAccessKey(id)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to delete access key")
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Access key was deleted successfully")
	}
}
