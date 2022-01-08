package status

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
	data, response, err := client.GetAllowlistingStatus()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
