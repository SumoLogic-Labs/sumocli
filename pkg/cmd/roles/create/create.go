package create

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdRoleCreate(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
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
			createRole(client, name, description, filter, users, capabilities, autofill, log)
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
	capabilities []string, autofill bool, log *zerolog.Logger) {
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
	apiResponse, httpResponse, errorResponse := client.CreateRole(body)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to create role")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
