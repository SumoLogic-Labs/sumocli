package get_capacity_info

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdFieldManagementGetCapacityInfo(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-capacity-info",
		Short: "Every account has a limited number of fields available." +
			"This command returns your account limitations and remaining quota",
		Run: func(cmd *cobra.Command, args []string) {
			getCapacityInfo(client)
		},
	}
	return cmd
}

func getCapacityInfo(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetFieldQuota()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
