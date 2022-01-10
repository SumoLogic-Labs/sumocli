package list

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdTokensList(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all tokens in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			listTokens(client)
		},
	}
	return cmd
}

func listTokens(client *cip.APIClient) {
	data, response, err := client.ListTokens()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
