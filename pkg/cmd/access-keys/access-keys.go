package access_keys

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	cmdAccessKeysCreate "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/create"
	cmdAccessKeysDelete "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/delete"
	cmdAccessKeysListAll "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/list-all"
	cmdAccessKeysListPersonal "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/list-personal"
	cmdAccessKeysUpdate "github.com/wizedkyle/sumocli/pkg/cmd/access-keys/update"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccessKeys(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access-keys",
		Short: "Manage access keys",
		Long:  "Commands that allow you to manage access keys in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdAccessKeysCreate.NewCmdAccessKeysCreate(client, log))
	cmd.AddCommand(cmdAccessKeysDelete.NewCmdAccessKeysDelete(client, log))
	cmd.AddCommand(cmdAccessKeysListAll.NewCmdAccessKeysListAll(client, log))
	cmd.AddCommand(cmdAccessKeysListPersonal.NewCmdAccessKeysListPersonal(client, log))
	cmd.AddCommand(cmdAccessKeysUpdate.NewCmdAccessKeysUpdate(client, log))
	return cmd
}
