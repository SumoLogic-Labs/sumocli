package get

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdCollectorGet() *cobra.Command {
	var ()

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic collector information",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			getCollector(log)
		},
	}
}

func getCollector(log zerolog.Logger) {

}
