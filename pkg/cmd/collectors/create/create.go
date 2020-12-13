package create

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdCollectorCreate() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic Hosted Collector",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			createCollector(log)
		},
	}

	return cmd
}

func createCollector(log zerolog.Logger) {

}
