package tokens

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	cmdTokensCreate "github.com/wizedkyle/sumocli/pkg/cmd/tokens/create"
	cmdTokensDelete "github.com/wizedkyle/sumocli/pkg/cmd/tokens/delete"
	cmdTokensGet "github.com/wizedkyle/sumocli/pkg/cmd/tokens/get"
	cmdTokensList "github.com/wizedkyle/sumocli/pkg/cmd/tokens/list"
	cmdTokensUpdate "github.com/wizedkyle/sumocli/pkg/cmd/tokens/update"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdTokens(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "Manage tokens",
		Long:  "Commands that allow you to manage Tokens in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdTokensCreate.NewCmdTokensCreate(client, log))
	cmd.AddCommand(cmdTokensDelete.NewCmdTokensDelete(client, log))
	cmd.AddCommand(cmdTokensGet.NewCmdTokensGet(client, log))
	cmd.AddCommand(cmdTokensList.NewCmdTokensList(client, log))
	cmd.AddCommand(cmdTokensUpdate.NewCmdTokensUpdate(client, log))
	return cmd
}
