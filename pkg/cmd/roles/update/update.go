package update

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"os"
)

func NewCmdRoleUpdate() *cobra.Command {
	var (
		id              string
		name            string
		description     string
		filterPredicate string
		users           []string
		capabilities    []string
		autofill        bool
		append          bool
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			updateRole(id, name, description, filterPredicate, users, capabilities, autofill, append)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to update")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name for the role")
	cmd.Flags().StringVar(&description, "description", "", "Specify the role description")
	cmd.Flags().StringVar(&filterPredicate, "filter", "", "Search filter for the role")
	cmd.Flags().StringSliceVar(&users, "users", []string{}, "Comma deliminated list of user ids to add to the role")
	cmd.Flags().StringSliceVar(&capabilities, "capabilities", []string{}, "Comma deliminated list of capabilities")
	cmd.Flags().BoolVar(&autofill, "autofill", true, "Is set to true by default.")
	cmd.Flags().BoolVar(&append, "append", true, "Is set to true by default, if set to false it will overwrite the role")

	return cmd
}

func updateRole(id string, name string, description string, filterPredicate string, users []string, capabilities []string, autofill bool, append bool) {
	if id == "" {
		fmt.Println("--id field needs to be set.")
		os.Exit(0)
	}

	var roleInfo api.RoleData
	requestUrl := "v1/roles/" + id

}
