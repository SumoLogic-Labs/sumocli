package reset_password

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdUserResetPassword(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "reset-password",
		Short: "Initiates a password reset for a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			userResetPassword(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user which requires a password reset.")
	cmd.MarkFlagRequired("id")
	return cmd
}

func userResetPassword(id string, client *cip.APIClient) {
	response, err := client.ResetPassword(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "User's password was reset successfully.")
	}
}
