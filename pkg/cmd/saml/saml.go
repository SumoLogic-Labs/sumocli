package saml

import (
	"github.com/spf13/cobra"
	NewCmdSamlAddAllowListUser "github.com/wizedkyle/sumocli/pkg/cmd/saml/add-allowlist-user"
	NewCmdSamlCreateConfiguration "github.com/wizedkyle/sumocli/pkg/cmd/saml/create-configuration"
	NewCmdSamlDeleteConfiguration "github.com/wizedkyle/sumocli/pkg/cmd/saml/delete-configuration"
	NewCmdSamlDisableLockdown "github.com/wizedkyle/sumocli/pkg/cmd/saml/disable-lockdown"
	NewCmdSamlEnableLockdown "github.com/wizedkyle/sumocli/pkg/cmd/saml/enable-lockdown"
	NewCmdSamlGetAllowListUsers "github.com/wizedkyle/sumocli/pkg/cmd/saml/get-allowlist-users"
	NewCmdSamlGetConfigurations "github.com/wizedkyle/sumocli/pkg/cmd/saml/get-configurations"
	NewCmdSamlRemoveAllowListUser "github.com/wizedkyle/sumocli/pkg/cmd/saml/remove-allowlist-user"
	NewCmdSamlUpdateConfiguration "github.com/wizedkyle/sumocli/pkg/cmd/saml/update-configuration"
)

func NewCmdSaml() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "saml",
		Short: "Manage SAML configuration",
		Long:  "Commands that allow you to manage the SAML configurations in your Sumo Logic tenant",
	}
	cmd.AddCommand(NewCmdSamlAddAllowListUser.NewCmdSamlAddAllowListUser())
	cmd.AddCommand(NewCmdSamlCreateConfiguration.NewCmdSamlCreateConfiguration())
	cmd.AddCommand(NewCmdSamlDeleteConfiguration.NewCmdSamlDeleteConfiguration())
	cmd.AddCommand(NewCmdSamlDisableLockdown.NewCmdSamlDisableLockdown())
	cmd.AddCommand(NewCmdSamlEnableLockdown.NewCmdSamlEnableLockdown())
	cmd.AddCommand(NewCmdSamlGetAllowListUsers.NewCmdSamlGetAllowListUsers())
	cmd.AddCommand(NewCmdSamlGetConfigurations.NewCmdSamlGetConfigurations())
	cmd.AddCommand(NewCmdSamlRemoveAllowListUser.NewCmdSamlRemoveAllowListUser())
	cmd.AddCommand(NewCmdSamlUpdateConfiguration.NewCmdSamlUpdateConfiguration())
	return cmd
}
