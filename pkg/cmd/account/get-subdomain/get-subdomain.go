package get_subdomain

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccountGetSubdomain(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-subdomain",
		Short: "Get the configured subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			getSubdomain(client)
		},
	}
	return cmd
}

func getSubdomain(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetSubdomain()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
