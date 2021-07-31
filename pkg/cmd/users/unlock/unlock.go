package unlock

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdUnlockUser(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "unlock",
		Short: "Unlocks a Sumo Logic user account",
		Run: func(cmd *cobra.Command, args []string) {
			unlockUser(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user account to unlock")
	cmd.MarkFlagRequired("id")
	return cmd
}

func unlockUser(id string, client *cip.APIClient) {
	httpResponse, errorResponse := client.UnlockUser(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "User's account was unlocked successfully.")
	}
}
