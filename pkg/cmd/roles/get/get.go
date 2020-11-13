package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/roles"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
)

func NewCmdRoleGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			getRoles(id)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to get")

	return cmd
}

func getRoles(id string) {
	var roleInfo roles.Role
	client := util2.GetHttpClient()
}
