package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdCollectorList(client *cip.APIClient) *cobra.Command {
	var (
		aliveBeforeDays int32
		filter          string
		limit           int32
		offset          int32
		offline         bool
	)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic collectors",
		Run: func(cmd *cobra.Command, args []string) {
			listCollectors(aliveBeforeDays, filter, limit, offset, offline, client)
		},
	}
	cmd.Flags().Int32Var(&aliveBeforeDays, "aliveBeforeDays", 100, "Minimum number of days the collectors have been offline (only used when offline is set to true)")
	cmd.Flags().StringVar(&filter, "filter", "", "Filters the collectors returned using either installed, hosted, dead or alive")
	cmd.Flags().Int32Var(&limit, "limit", 1000, "Maximum number of collectors returned")
	cmd.Flags().Int32Var(&offset, "offset", 0, "Offset into the list of collectors")
	cmd.Flags().BoolVar(&offline, "offline", false, "Lists offline collectors")
	return cmd
}

func listCollectors(aliveBeforeDays int32, filter string, limit int32, offset int32, offline bool, client *cip.APIClient) {
	if offline == true {
		apiResponse, httpResponse, errorResponse := client.ListOfflineCollectors(&types.ListCollectorsOfflineOpts{
			AliveBeforeDays: optional.NewInt32(aliveBeforeDays),
			Limit:           optional.NewInt32(limit),
			Offset:          optional.NewInt32(offset),
		})
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse, errorResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if offline == false {
		apiResponse, httpResponse, errorResponse := client.ListCollectors(&types.ListCollectorsOpts{
			Filter: optional.NewString(filter),
			Limit:  optional.NewInt32(limit),
			Offset: optional.NewInt32(offset),
		})
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse, errorResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	}
}
