package get_capacity_info

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
	data, response, err := client.GetFieldQuota()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
