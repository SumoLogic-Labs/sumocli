package create

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewCmdRoleCreate(client *cip.APIClient) *cobra.Command {
	var (
		name         string
		description  string
		filter       string
		users        []string
		capabilities []string
		autofill     bool
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			createRole(client, name, description, filter, users, capabilities, autofill)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Name of the role.")
	cmd.Flags().StringVar(&description, "description", "", "Description for the role.")
	cmd.Flags().StringVar(&filter, "filter", "", "Search filter for the role.")
	cmd.Flags().StringSliceVar(&users, "users", []string{}, "Comma deliminated list of user ids to add to the role.")
	cmd.Flags().StringSliceVar(&capabilities, "capabilities", []string{}, "Comma deliminated list of capabilities.")
	cmd.Flags().BoolVar(&autofill, "autofill", true, "Is set to true by default.")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createRole(client *cip.APIClient, name string, description string, filter string, users []string,
	capabilities []string, autofill bool) {
	for i, capability := range capabilities {
		if cmdutils.ValidateCapabilities(capability) == false {
			log.Error().Msg(capability + " is not a valid Sumo Logic role capability.")
		}
		i++
	}
	body := types.CreateRoleDefinition{
		Name:                 name,
		Description:          description,
		FilterPredicate:      filter,
		Users:                users,
		Capabilities:         capabilities,
		AutofillDependencies: autofill,
	}
	data, response, err := client.CreateRole(body)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
