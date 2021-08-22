package get_upgradable_collectors

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdGetUpgradableCollectors(client *cip.APIClient) *cobra.Command {
	var (
		toVersion string
		offset    int32
		limit     int32
	)
	cmd := &cobra.Command{
		Use:   "get-upgradable-collectors",
		Short: "Gets collectors in Sumo Logic that are upgradable",
		Run: func(cmd *cobra.Command, args []string) {
			getUpgradableCollectors(toVersion, offset, limit, client)
		},
	}
	cmd.Flags().StringVar(&toVersion, "version", "", "Collector build to upgrade to, if not specified defaults to the latest version")
	cmd.Flags().Int32Var(&offset, "offset", 0, "Offset into the list of collectors")
	cmd.Flags().Int32Var(&limit, "limit", 50, "Maximum number of collectors to return")
	return cmd
}

func getUpgradableCollectors(toVersion string, offset int32, limit int32, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetUpgradableCollectors(&types.GetUpgradableCollectorsOpts{
		Limit:     optional.NewInt32(limit),
		Offset:    optional.NewInt32(offset),
		ToVersion: optional.NewString(toVersion),
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
