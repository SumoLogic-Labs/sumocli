package password_policy

import (
	NewCmdPasswordPolicyGet "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/password-policy/get"
	NewCmdPasswordPolicyUpdate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/password-policy/update"
	"github.com/spf13/cobra"
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
