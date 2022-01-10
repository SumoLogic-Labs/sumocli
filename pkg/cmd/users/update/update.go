package update

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdUserUpdate(client *cip.APIClient) *cobra.Command {
	var (
		id        string
		firstName string
		lastName  string
		isActive  bool
		roleIds   []string
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			updateUser(id, firstName, lastName, isActive, roleIds, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user to update.")
	cmd.Flags().StringVar(&firstName, "firstName", "", "First name of the user.")
	cmd.Flags().StringVar(&lastName, "lastName", "", "Last name for the user.")
	cmd.Flags().BoolVar(&isActive, "isActive", true, "True if the account is active, false if it is deactivated")
	cmd.Flags().StringSliceVar(&roleIds, "roleIds", []string{}, "Comma deliminated list of Role Ids.")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("firstName")
	cmd.MarkFlagRequired("lastName")
	cmd.MarkFlagRequired("isActive")
	cmd.MarkFlagRequired("roleIds")
	return cmd
}

func updateUser(id string, firstName string, lastName string, isActive bool, roleIds []string, client *cip.APIClient) {
	data, response, err := client.UpdateUser(types.UpdateUserDefinition{
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  isActive,
		RoleIds:   roleIds,
	},
		id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
