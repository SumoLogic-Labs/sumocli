package live_tail

import (
	cmdLiveTailStart "github.com/SumoLogic-Labs/sumocli/pkg/cmd/live-tail/start"
	"github.com/spf13/cobra"
)

func NewCmdLiveTail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "live-tail",
		Short: "Manages access to live tail via CLI",
	}

	cmd.AddCommand(cmdLiveTailStart.StartLiveTailCmd())
	return cmd
}
