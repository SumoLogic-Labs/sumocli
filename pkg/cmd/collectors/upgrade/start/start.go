package start

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
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
	apiResponse, httpResponse, errorResponse := client.CreateUpgradeOrDowngradeTask(types.CreateUpgradeOrDowngradeRequest{
		CollectorId: idInt,
		ToVersion:   toVersion,
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
