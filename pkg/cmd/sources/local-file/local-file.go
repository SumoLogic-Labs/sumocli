package local_file

import (
	"github.com/spf13/cobra"
	NewCmdLocalFileSourceCreate "github.com/wizedkyle/sumocli/pkg/cmd/sources/local-file/create"
)

func NewCmdLocalFileSources() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "local-file <command>",
		Short: "Manage local file sources",
	}
	cmd.AddCommand(NewCmdLocalFileSourceCreate.NewCmdCreateLocalFileSource())
	return cmd
}
