package roles

import (
	"github.com/spf13/cobra"
	cmdRoleList "github.com/wizedkyle/sumocli/pkg/cmd/roles/list"
)

func NewCmdRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "role <command>",
		Short: "",
		Long:  "",
		//Example: heredoc.Doc(),
	}

	cmd.AddCommand(cmdRoleList.NewCmdRoleList())
	return cmd
}
