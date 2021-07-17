package list

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdTokensList(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all tokens in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			listTokens(client, log)
		},
	}
	return cmd
}

func listTokens(client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.ListTokens()
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to list tokens")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
