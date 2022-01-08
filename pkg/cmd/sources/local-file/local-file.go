package local_file

import (
	NewCmdLocalFileSourceCreate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/sources/local-file/create"
	"github.com/spf13/cobra"
)

func NewCmdLocalFileSources() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "local-file <command>",
		Short: "Manage local file sources",
	}
	cmd.AddCommand(NewCmdLocalFileSourceCreate.NewCmdCreateLocalFileSource())
	return cmd
}
