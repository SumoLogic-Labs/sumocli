package reset

import (
	"github.com/spf13/cobra"
)

func NewCmdUserResetPassword() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.PersistentFlags().StringVar(&id, "id", "", "Specify the id of the user which requires a password reset.")

	return cmd
}

func userResetPassword() {
	// https://api.au.sumologic.com/docs/#operation/resetPassword
}
