package start

import (
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
	"strconv"
)

func NewCmdUpgradeStart(client *cip.APIClient) *cobra.Command {
	var (
		id        string
		toVersion string
	)
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts an upgrade or downgrade of an existing installed collector",
		Run: func(cmd *cobra.Command, args []string) {
			upgradeStart(id, toVersion, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Id of the collector to upgrade")
	cmd.Flags().StringVar(&toVersion, "version", "", "Version to upgrade or downgrade the collector to")
	cmd.MarkFlagRequired("collectorId")
	return cmd
}

func upgradeStart(id string, toVersion string, client *cip.APIClient) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("failed to convert string to int")
	}
	data, response, err := client.CreateUpgradeOrDowngradeTask(types.CreateUpgradeOrDowngradeRequest{
		CollectorId: idInt,
		ToVersion:   toVersion,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
