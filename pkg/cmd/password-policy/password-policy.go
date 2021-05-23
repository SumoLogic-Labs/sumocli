package password_policy

import (
	"github.com/spf13/cobra"
	NewCmdPasswordPolicyGet "github.com/wizedkyle/sumocli/pkg/cmd/password-policy/get"
	NewCmdPasswordPolicyUpdate "github.com/wizedkyle/sumocli/pkg/cmd/password-policy/update"
)

func NewCmdPasswordPolicy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password-policy",
		Short: "Manage password policy",
		Long:  "Commands that allow you to modify the password policy in your Sumo Logic tenant",
	}
	cmd.AddCommand(NewCmdPasswordPolicyGet.NewCmdPasswordPolicyGet())
	cmd.AddCommand(NewCmdPasswordPolicyUpdate.NewCmdPasswordPolicyUpdate())
	return cmd
}
