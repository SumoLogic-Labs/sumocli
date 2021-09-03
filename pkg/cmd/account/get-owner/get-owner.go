package get_owner

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAccountGetOwner(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-owner",
		Short: "Returns the user identifier as the account owner.",
		Run: func(cmd *cobra.Command, args []string) {
			getOwner(client)
		},
	}
	return cmd
}

func getOwner(client *cip.APIClient) {
	userId, httpResponse, errorResponse := client.GetAccountOwner()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		userResponse, httpResponse, errorResponse := client.GetUser(userId)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse, errorResponse)
		} else {
			cmdutils.Output(userResponse, httpResponse, errorResponse, "")
		}
	}
}
