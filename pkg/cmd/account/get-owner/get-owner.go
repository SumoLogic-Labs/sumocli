package get_owner

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccountGetOwner(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-owner",
		Short: "Returns the user identifier as the account owner.",
		Run: func(cmd *cobra.Command, args []string) {
			getOwner(client, log)
		},
	}
	return cmd
}

func getOwner(client *cip.APIClient, log *zerolog.Logger) {
	userId, _, errorResponse := client.GetAccountOwner()
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get account owner")
	} else {
		userResponse, httpResponse, errorResponse := client.GetUser(userId)
		if errorResponse != nil {
			log.Error().Err(errorResponse).Msg("failed to get user")
		} else {
			cmdutils.Output(userResponse, httpResponse, errorResponse, "")
		}
	}
}
