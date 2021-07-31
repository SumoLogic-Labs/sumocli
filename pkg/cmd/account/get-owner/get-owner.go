package get_owner

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
		cmdutils.OutputError(httpResponse)
	} else {
		userResponse, httpResponse, errorResponse := client.GetUser(userId)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(userResponse, httpResponse, errorResponse, "")
		}
	}
}
