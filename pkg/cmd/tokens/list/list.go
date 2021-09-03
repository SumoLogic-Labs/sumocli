package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdTokensList(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all tokens in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			listTokens(client)
		},
	}
	return cmd
}

func listTokens(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.ListTokens()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
