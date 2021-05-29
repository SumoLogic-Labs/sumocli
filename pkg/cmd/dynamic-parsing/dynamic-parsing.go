package dynamic_parsing

import (
	"github.com/spf13/cobra"
	NewCmdDynamicParsingCreate "github.com/wizedkyle/sumocli/pkg/cmd/dynamic-parsing/create"
	NewCmdDynamicParsingDelete "github.com/wizedkyle/sumocli/pkg/cmd/dynamic-parsing/delete"
	NewCmdDynamicParsingGet "github.com/wizedkyle/sumocli/pkg/cmd/dynamic-parsing/get"
	NewCmdDynamicParsingList "github.com/wizedkyle/sumocli/pkg/cmd/dynamic-parsing/list"
	NewCmdDynamicParsingUpdate "github.com/wizedkyle/sumocli/pkg/cmd/dynamic-parsing/update"
)

func NewCmdDynamicParsing() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dynamic-parsing",
		Short: "Manage dynamic parsing settings",
		Long:  "Dynamic Parsing allows automatic field extraction from your log messages when you run a search.",
	}
	cmd.AddCommand(NewCmdDynamicParsingCreate.NewCmdDynamicParsingCreate())
	cmd.AddCommand(NewCmdDynamicParsingDelete.NewCmdDynamicParsingDelete())
	cmd.AddCommand(NewCmdDynamicParsingGet.NewCmdDynamicParsingGet())
	cmd.AddCommand(NewCmdDynamicParsingList.NewCmdDynamicParsingList())
	cmd.AddCommand(NewCmdDynamicParsingUpdate.NewCmdDynamicParsingUpdate())
	return cmd
}
