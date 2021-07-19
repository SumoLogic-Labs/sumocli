package account

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	NewCmdAccountCreateSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/create-subdomain"
	NewCmdAccountDeleteSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/delete-subdomain"
	NewCmdAccountGetOwner "github.com/wizedkyle/sumocli/pkg/cmd/account/get-owner"
	NewCmdAccountGetSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/get-subdomain"
	NewCmdAccountRecoverSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/recover-subdomain"
	NewCmdAccountUpdateSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/update-subdomain"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccount(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "Manage account settings",
		Long:  "Commands that allow you to manage your Sumo Logic account settings",
	}
	cmd.AddCommand(NewCmdAccountCreateSubdomain.NewCmdAccountCreateSubdomain(client, log))
	cmd.AddCommand(NewCmdAccountDeleteSubdomain.NewCmdAccountDeleteSubdomain(client, log))
	cmd.AddCommand(NewCmdAccountGetOwner.NewCmdAccountGetOwner(client, log))
	cmd.AddCommand(NewCmdAccountGetSubdomain.NewCmdAccountGetSubdomain(client, log))
	cmd.AddCommand(NewCmdAccountRecoverSubdomain.NewCmdAccountRecoverSubdomain(client, log))
	cmd.AddCommand(NewCmdAccountUpdateSubdomain.NewCmdAccountUpdateSubdomain(client, log))
	return cmd
}
