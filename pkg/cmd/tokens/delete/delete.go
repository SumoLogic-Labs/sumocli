package delete

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdTokensDelete(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a token with the given identifier in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteToken(id, client, log)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify id of the token to delete")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteToken(id string, client *cip.APIClient, log *zerolog.Logger) {
	httpResponse, errorResponse := client.DeleteToken(id)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to delete token")
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Token was deleted.")
	}
}
