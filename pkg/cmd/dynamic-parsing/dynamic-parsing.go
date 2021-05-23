package dynamic_parsing

import (
	"github.com/spf13/cobra"
	NewCmdDynamicParsingList "github.com/wizedkyle/sumocli/pkg/cmd/dynamic-parsing/list"
)

func NewCmdDynamicParsing() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dynamic-parsing",
		Short: "Manage dynamic parsing settings",
		Long:  "Dynamic Parsing allows automatic field extraction from your log messages when you run a search.",
	}
	cmd.AddCommand(NewCmdDynamicParsingList.NewCmdDynamicParsingList())
	return cmd
}
