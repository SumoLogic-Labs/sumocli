package update

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdRoleUpdate(client *cip.APIClient) *cobra.Command {
	var (
		id           string
		name         string
		description  string
		filter       string
		users        []string
		capabilities []string
		autofill     bool
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates a Sumo Logic role.",
		Run: func(cmd *cobra.Command, args []string) {
			updateRole(client, id, name, description, filter, users, capabilities, autofill)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to update.")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name for the role.")
	cmd.Flags().StringVar(&description, "description", "", "Specify the role description.")
	cmd.Flags().StringVar(&filter, "filter", "", "Search filter for the role.")
	cmd.Flags().StringSliceVar(&users, "users", []string{}, "Comma deliminated list of user ids to add to the role.")
	cmd.Flags().StringSliceVar(&capabilities, "capabilities", []string{}, "Comma deliminated list of capabilities.")
	cmd.Flags().BoolVar(&autofill, "autofill", true, "Is set to true by default.")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	return cmd
}

func updateRole(client *cip.APIClient, id string, name string, description string, filter string, users []string,
	capabilities []string, autofill bool) {
	for i, capability := range capabilities {
		if cmdutils.ValidateCapabilities(capability) == false {
			log.Error().Msg(capability + " is not a valid Sumo Logic role capability.")
		}
		i++
	}
	body := types.UpdateRoleDefinition{
		Name:                 name,
		Description:          description,
		FilterPredicate:      filter,
		Users:                users,
		Capabilities:         capabilities,
		AutofillDependencies: autofill,
	}
	apiResponse, httpResponse, errorResponse := client.UpdateRole(body, id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
