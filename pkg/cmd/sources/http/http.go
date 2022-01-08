package http

import (
	NewCmdCreateHttpSource "github.com/SumoLogic-Labs/sumocli/pkg/cmd/sources/http/create"
	NewCmdUpdateHttpSource "github.com/SumoLogic-Labs/sumocli/pkg/cmd/sources/http/update"
	"github.com/spf13/cobra"
)

func NewCmdHttpSources() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http <command>",
		Short: "Manage https sources",
	}
	cmd.AddCommand(NewCmdCreateHttpSource.NewCmdCreateHttpSource())
	cmd.AddCommand(NewCmdUpdateHttpSource.NewCmdUpdateHttpSource())
	return cmd
}
