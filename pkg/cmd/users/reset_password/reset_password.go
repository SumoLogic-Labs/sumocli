package reset_password

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdUserResetPassword(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "reset-password",
		Short: "Initiates a password reset for a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			userResetPassword(id, client, log)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user which requires a password reset.")
	cmd.MarkFlagRequired("id")
	return cmd
}

func userResetPassword(id string, client *cip.APIClient, log *zerolog.Logger) {
	httpResponse, errorResponse := client.ResetPassword(id)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to reset password")
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "User's password was reset successfully.")
	}
}
