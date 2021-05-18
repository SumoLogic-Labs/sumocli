package update

import (
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgetsUpdate() *cobra.Command {
	var ()

	cmd := &cobra.Command{
		Use:   "update",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}

func updateIngestBudget() {

}
