package http

import (
	"github.com/spf13/cobra"
	NewCmdCreateHttpSource "github.com/wizedkyle/sumocli/pkg/cmd/sources/http/create"
)

func NewCmdHttpSources() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http <command>",
		Short: "Manage https sources",
	}
	cmd.AddCommand(NewCmdCreateHttpSource.NewCmdCreateHttpSource())
	return cmd
}
