package live_tail

import (
	"github.com/spf13/cobra"
	cmdLiveTailStart "github.com/wizedkyle/sumocli/pkg/cmd/live-tail/start"
)

func NewCmdLiveTail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "live tail",
		Short: "Manages access to live tail via CLI",
	}

	cmd.AddCommand(cmdLiveTailStart.StartLiveTailCmd())
	return cmd
}
