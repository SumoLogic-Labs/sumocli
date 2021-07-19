package create

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdUserCreate(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
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
			user(firstName, lastName, emailAddress, roleIds, client, log)
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

func user(firstName string, lastName string, emailAddress string, roleIds []string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.CreateUser(types.CreateUserDefinition{
		FirstName: firstName,
		LastName:  lastName,
		Email:     emailAddress,
		RoleIds:   roleIds,
	})
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to create user")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
