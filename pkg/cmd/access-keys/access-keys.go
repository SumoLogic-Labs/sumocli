package access_keys

import (
	"github.com/spf13/cobra"
	cmdAccessKeysCreate "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/create"
	cmdAccessKeysDelete "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/delete"
	cmdAccessKeysListAll "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/list_all"
	cmdAccessKeysListPersonal "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/list_personal"
	cmdAccessKeysUpdate "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/update"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccessKeys(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access-keys",
		Short: "Manage access keys",
		Long:  "Commands that allow you to manage access keys in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdAccessKeysCreate.NewCmdAccessKeysCreate(client))
	cmd.AddCommand(cmdAccessKeysDelete.NewCmdAccessKeysDelete(client))
	cmd.AddCommand(cmdAccessKeysListAll.NewCmdAccessKeysListAll(client))
	cmd.AddCommand(cmdAccessKeysListPersonal.NewCmdAccessKeysListPersonal(client))
	cmd.AddCommand(cmdAccessKeysUpdate.NewCmdAccessKeysUpdate(client))
	return cmd
}
