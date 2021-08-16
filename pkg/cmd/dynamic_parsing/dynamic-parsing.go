package dynamic_parsing

import (
	"github.com/spf13/cobra"
	NewCmdDynamicParsingCreate "github.com/wizedkyle/sumocli/pkg/cmd/dynamic_parsing/create"
	NewCmdDynamicParsingDelete "github.com/wizedkyle/sumocli/pkg/cmd/dynamic_parsing/delete"
	NewCmdDynamicParsingGet "github.com/wizedkyle/sumocli/pkg/cmd/dynamic_parsing/get"
	NewCmdDynamicParsingList "github.com/wizedkyle/sumocli/pkg/cmd/dynamic_parsing/list"
	NewCmdDynamicParsingUpdate "github.com/wizedkyle/sumocli/pkg/cmd/dynamic_parsing/update"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdDynamicParsing(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dynamic-parsing",
		Short: "Manage dynamic parsing rules",
		Long:  "Dynamic Parsing allows automatic field extraction from your log messages when you run a search.",
	}
	cmd.AddCommand(NewCmdDynamicParsingCreate.NewCmdDynamicParsingCreate(client))
	cmd.AddCommand(NewCmdDynamicParsingDelete.NewCmdDynamicParsingDelete(client))
	cmd.AddCommand(NewCmdDynamicParsingGet.NewCmdDynamicParsingGet(client))
	cmd.AddCommand(NewCmdDynamicParsingList.NewCmdDynamicParsingList(client))
	cmd.AddCommand(NewCmdDynamicParsingUpdate.NewCmdDynamicParsingUpdate(client))
	return cmd
}
