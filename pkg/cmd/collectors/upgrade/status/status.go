package status

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdUpgradableCollectorStatus(client *cip.APIClient) *cobra.Command {
	var upgradeTaskId string
	cmd := &cobra.Command{
		Use: "status",
		Long: `Gets the status of a collector upgrade or downgrade.
The status of the upgrade can be one of the following
0 - not started
1 - pending, the upgrade is issued waiting a response from the Collector
2 - succeeded
3 - failed
6 - progressing, the upgrade is running on the Collector`,
		Run: func(cmd *cobra.Command, args []string) {
			upgradableCollectorStatus(upgradeTaskId, client)
		},
	}
	cmd.Flags().StringVar(&upgradeTaskId, "upgradeTaskId", "", "Id to the upgrade task")
	cmd.MarkFlagRequired("upgradeTaskId")
	return cmd
}

func upgradableCollectorStatus(upgradeTaskId string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetUpgradeOrDowngradeTaskStatus(upgradeTaskId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
