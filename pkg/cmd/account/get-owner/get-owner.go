package get_owner

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAccountGetOwner(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-owner",
		Short: "Returns the user identifier as the account owner.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getOwner(client)
		},
	}
	return cmd
}

func getOwner(client *cip.APIClient) {
	userId, response, err := client.GetAccountOwner()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		userResponse, response, err := client.GetUser(userId)
		if err != nil {
			cmdutils.OutputError(response, err)
		} else {
			cmdutils.Output(userResponse, response, err, "")
		}
	}
}
