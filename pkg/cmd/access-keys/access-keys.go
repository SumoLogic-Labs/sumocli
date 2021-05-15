package access_keys

import (
	"github.com/spf13/cobra"
	cmdAccessKeysCreate "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/create"
	cmdAccessKeysDelete "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/delete"
	cmdAccessKeysListAll "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/list-all"
	cmdAccessKeysListPersonal "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/list-personal"
	cmdAccessKeysUpdate "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/update"
)

func NewCmdAccessKeys() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access-keys",
		Short: "Managed access keys",
		Long:  "Commands that allow you to manage access keys in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdAccessKeysCreate.NewCmdAccessKeysCreate())
	cmd.AddCommand(cmdAccessKeysDelete.NewCmdAccessKeysDelete())
	cmd.AddCommand(cmdAccessKeysListAll.NewCmdAccessKeysListAll())
	cmd.AddCommand(cmdAccessKeysListPersonal.NewCmdAccessKeysListPersonal())
	cmd.AddCommand(cmdAccessKeysUpdate.NewCmdAccessKeysUpdate())
	return cmd
}
