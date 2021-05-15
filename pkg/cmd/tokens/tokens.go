package tokens

import (
	"github.com/spf13/cobra"
	cmdTokensCreate "github.com/wizedkyle/sumocli/pkg/cmd/tokens/create"
	cmdTokensDelete "github.com/wizedkyle/sumocli/pkg/cmd/tokens/delete"
	cmdTokensGet "github.com/wizedkyle/sumocli/pkg/cmd/tokens/get"
	cmdTokensList "github.com/wizedkyle/sumocli/pkg/cmd/tokens/list"
	cmdTokensUpdate "github.com/wizedkyle/sumocli/pkg/cmd/tokens/update"
)

func NewCmdTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "Manage tokens",
		Long:  "Commands that allow you to manage Tokens in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdTokensCreate.NewCmdTokensCreate())
	cmd.AddCommand(cmdTokensDelete.NewCmdTokensDelete())
	cmd.AddCommand(cmdTokensGet.NewCmdTokensGet())
	cmd.AddCommand(cmdTokensList.NewCmdTokensList())
	cmd.AddCommand(cmdTokensUpdate.NewCmdTokensUpdate())
	return cmd
}
