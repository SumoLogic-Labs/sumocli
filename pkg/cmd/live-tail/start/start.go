package start

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func StartLiveTailCmd() *cobra.Command {
	var (
		tailId string
	)
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a live tail session",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			startLiveTailSession(tailId, log)
		},
	}

	cmd.Flags().StringVar(&tailId, "tailId", "", "Test argument")
	return cmd
}

func createLiveTailSession() {

}

func startLiveTailSession(tailId string, log zerolog.Logger) {

}
