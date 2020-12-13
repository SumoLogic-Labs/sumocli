package list

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdCollectorList() *cobra.Command {
	var (
		filter  string
		limit   int
		offset  int
		offline bool
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic collectors",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			listCollectors(offline, log)
		},
	}

	cmd.Flags().StringVar(&filter, "filter", "", "Filters the collectors returned using either installed, hosted, dead or alive")
	cmd.Flags().IntVar(&limit, "limit", 1000, "Maximum number of collectors returned")
	cmd.Flags().IntVar(&offset, "offset", 0, "Offset into the list of collectors")
	cmd.Flags().BoolVar(&offline, "offline", false, "Lists offline collectors")
	return cmd
}

func listCollectors(offline bool, log zerolog.Logger) {
	if offline == true {

	}
}
