package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdTokensGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a token with the given identifier in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			getToken(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the token to retrieve")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getToken(id string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetToken(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
