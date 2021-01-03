package azure

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/internal/config"
	cmdAzureCreate "github.com/wizedkyle/sumocli/pkg/cmd/azure/create"
)

func NewCmdAzure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "azure <command>",
		Short: "Creates and deletes Azure infrastructure for log and metric collection",
	}

	config.AddAzureFlags(cmd)

	cmd.AddCommand(cmdAzureCreate.NewCmdAzureCreate())
	return cmd
}
