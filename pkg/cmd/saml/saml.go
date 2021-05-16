package saml

import (
	"github.com/spf13/cobra"
	NewCmdSamlGetConfigurations "github.com/wizedkyle/sumocli/pkg/cmd/saml/get-configurations"
)

func NewCmdSaml() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "saml",
		Short: "Manage SAML configuration",
		Long:  "Commands that allow you to manage the SAML configurations in your Sumo Logic tenant",
	}
	cmd.AddCommand(NewCmdSamlGetConfigurations.NewCmdSamlGetConfigurations())
	return cmd
}
