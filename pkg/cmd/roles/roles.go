package roles

import (
	"github.com/spf13/cobra"
	cmdRoleList "github.com/wizedkyle/sumocli/pkg/cmd/roles/list"
)

type Role struct {
	Data []RoleData `json:"data"`
}

type RoleData struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	FilterPredicate      string   `json:"filterPredicate"`
	Users                []string `json:"users"`
	Capabilities         []string `json:"capabilities"`
	AutofillDependencies bool     `json:"autofillDependencies"`
	CreatedAt            string   `json:"createdAt"`
	CreatedBy            string   `json:"createdBy"`
	ModifiedAt           string   `json:"modifiedAt"`
	ModifiedBy           string   `json:"modifiedBy"`
	Id                   string   `json:"id"`
	SystemDefined        bool     `json:"systemDefined"`
}

func NewCmdRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roles <command>",
		Short: "Manage roles",
		Long:  "Work with Sumo Logic roles",
	}

	cmd.AddCommand(cmdRoleList.NewCmdRoleList())
	return cmd
}
