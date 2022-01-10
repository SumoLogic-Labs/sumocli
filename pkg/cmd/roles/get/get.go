package get

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdRoleGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic role information",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getRole(client, id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to get")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getRole(client *cip.APIClient, id string) {
	data, response, err := client.GetRole(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
