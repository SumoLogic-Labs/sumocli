package password_policy

import (
	NewCmdPasswordPolicyGet "github.com/SumoLogic-Labs/sumocli/pkg/cmd/password_policy/get"
	NewCmdPasswordPolicyUpdate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/password_policy/update"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdPasswordPolicy(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password-policy",
		Short: "Manage password policy",
		Long:  "Commands that allow you to modify the password policy in your Sumo Logic tenant",
	}
	cmd.AddCommand(NewCmdPasswordPolicyGet.NewCmdPasswordPolicyGet(client))
	cmd.AddCommand(NewCmdPasswordPolicyUpdate.NewCmdPasswordPolicyUpdate(client))
	return cmd
}
