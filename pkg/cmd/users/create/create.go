package create

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdUserCreate(client *cip.APIClient) *cobra.Command {
	var (
		firstName    string
		lastName     string
		emailAddress string
		roleIds      []string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic user account",
		Run: func(cmd *cobra.Command, args []string) {
			user(firstName, lastName, emailAddress, roleIds, client)
		},
	}
	cmd.Flags().StringVar(&firstName, "firstName", "", "First name of the user")
	cmd.Flags().StringVar(&lastName, "lastName", "", "Last name of the user")
	cmd.Flags().StringVar(&emailAddress, "email", "", "Email address of the user")
	cmd.Flags().StringSliceVar(&roleIds, "roleIds", []string{}, "Comma deliminated list of Role Ids")
	cmd.MarkFlagRequired("firstName")
	cmd.MarkFlagRequired("lastName")
	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("roleIds")
	return cmd
}

func user(firstName string, lastName string, emailAddress string, roleIds []string, client *cip.APIClient) {
	data, response, err := client.CreateUser(types.CreateUserDefinition{
		FirstName: firstName,
		LastName:  lastName,
		Email:     emailAddress,
		RoleIds:   roleIds,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
