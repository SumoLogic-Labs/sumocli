package unlock

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "User's account was unlocked successfully.")
	}
}
