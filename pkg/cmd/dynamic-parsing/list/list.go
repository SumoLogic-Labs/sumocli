package list

import (
	"github.com/spf13/cobra"
)

func NewCmdDynamicParsingList() *cobra.Command {
	var ()

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all dynamic parsing rules.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}

func listDynamicParsingRules() {

}
