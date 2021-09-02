package status

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdServiceAllowlistStatus(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status of the service allowlisting functionality for login/API authentication or content sharing for the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getServiceAllowlistStatus(client)
		},
	}
	return cmd
}

func getServiceAllowlistStatus(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetAllowlistingStatus()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
