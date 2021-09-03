package get_subdomain

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
