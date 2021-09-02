package reset_password

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdUserResetPassword(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "reset-password",
		Short: "Initiates a password reset for a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			userResetPassword(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user which requires a password reset.")
	cmd.MarkFlagRequired("id")
	return cmd
}

func userResetPassword(id string, client *cip.APIClient) {
	httpResponse, errorResponse := client.ResetPassword(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "User's password was reset successfully.")
	}
}
