package delete

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdTokensDelete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a token with the given identifier in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			deleteToken(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify id of the token to delete")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteToken(id string, client *cip.APIClient) {
	response, err := client.DeleteToken(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Token was deleted.")
	}
}
