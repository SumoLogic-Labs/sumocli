package update

import (
	"github.com/spf13/cobra"
)

func NewCmdCollectorUpdate() *cobra.Command {
	var ()

	cmd := &cobra.Command{
		Use:   "update",
		Short: "updates a Sumo Logic collector settings",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}

func updateCollector() {

}
