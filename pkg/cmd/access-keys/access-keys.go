package access_keys

import (
	"github.com/spf13/cobra"
	cmdAccessKeysListAll "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/list-all"
	cmdAccessKeysListPersonal "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/list-personal"
)

func NewCmdAccessKeys() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access-keys",
		Short: "Managed access keys",
		Long:  "Commands that allow you to manage access keys in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdAccessKeysListAll.NewCmdAccessKeysListAll())
	cmd.AddCommand(cmdAccessKeysListPersonal.NewCmdAccessKeysListPersonal())
	return cmd
}
