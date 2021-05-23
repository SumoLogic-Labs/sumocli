package account

import (
	"github.com/spf13/cobra"
	NewCmdAccountCreateSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/create-subdomain"
	NewCmdAccountDeleteSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/delete-subdomain"
	NewCmdAccountGetOwner "github.com/wizedkyle/sumocli/pkg/cmd/account/get-owner"
	NewCmdAccountGetSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/get-subdomain"
	NewCmdAccountRecoverSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/recover-subdomain"
	NewCmdAccountUpdateSubdomain "github.com/wizedkyle/sumocli/pkg/cmd/account/update-subdomain"
)

func NewCmdAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "Manage account settings",
		Long:  "Commands that allow you to manage your Sumo Logic account settings",
	}
	cmd.AddCommand(NewCmdAccountCreateSubdomain.NewCmdAccountCreateSubdomain())
	cmd.AddCommand(NewCmdAccountDeleteSubdomain.NewCmdAccountDeleteSubdomain())
	cmd.AddCommand(NewCmdAccountGetOwner.NewCmdAccountGetOwner())
	cmd.AddCommand(NewCmdAccountGetSubdomain.NewCmdAccountGetSubdomain())
	cmd.AddCommand(NewCmdAccountRecoverSubdomain.NewCmdAccountRecoverSubdomain())
	cmd.AddCommand(NewCmdAccountUpdateSubdomain.NewCmdAccountUpdateSubdomain())
	return cmd
}
