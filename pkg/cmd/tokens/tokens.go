package tokens

import (
	cmdTokensCreate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/tokens/create"
	cmdTokensDelete "github.com/SumoLogic-Labs/sumocli/pkg/cmd/tokens/delete"
	cmdTokensGet "github.com/SumoLogic-Labs/sumocli/pkg/cmd/tokens/get"
	cmdTokensList "github.com/SumoLogic-Labs/sumocli/pkg/cmd/tokens/list"
	cmdTokensUpdate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/tokens/update"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdTokens(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "Manage tokens",
		Long:  "Commands that allow you to manage Tokens in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdTokensCreate.NewCmdTokensCreate(client))
	cmd.AddCommand(cmdTokensDelete.NewCmdTokensDelete(client))
	cmd.AddCommand(cmdTokensGet.NewCmdTokensGet(client))
	cmd.AddCommand(cmdTokensList.NewCmdTokensList(client))
	cmd.AddCommand(cmdTokensUpdate.NewCmdTokensUpdate(client))
	return cmd
}
