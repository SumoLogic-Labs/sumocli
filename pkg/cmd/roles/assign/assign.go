package assign

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdRoleAssign(client *cip.APIClient) *cobra.Command {
	var (
		roleId string
		userId string
	)
	cmd := &cobra.Command{
		Use:   "assign",
		Short: "Assigns the specified Sumo Logic user to the role.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			assignRoleToUser(client, roleId, userId)
		},
	}
	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the identifier of the role to assign.")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the identifier of the user to assign the role to.")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func assignRoleToUser(client *cip.APIClient, roleId string, userId string) {
	authentication.ConfirmCredentialsSet(client)
	data, response, err := client.AssignRoleToUser(roleId, userId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
