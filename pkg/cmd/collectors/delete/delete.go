package delete

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdCollectorDelete(client *cip.APIClient) *cobra.Command {
	var (
		aliveBeforeDays int32
		id              string
		offline         bool
	)
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic collector",
		Run: func(cmd *cobra.Command, args []string) {
			deleteCollector(aliveBeforeDays, id, offline, client)
		},
	}
	cmd.Flags().Int32Var(&aliveBeforeDays, "aliveBeforeDays", 100, "Minimum number of days the collectors have been offline")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the collector to delete")
	cmd.Flags().BoolVar(&offline, "offline", false, "Removes all offline collectors")
	return cmd
}

func deleteCollector(aliveBeforeDays int32, id string, offline bool, client *cip.APIClient) {
	if id != "" {
		response, err := client.DeleteCollectorById(id)
		if err != nil {
			cmdutils.OutputError(response, err)
		} else {
			cmdutils.Output(nil, response, err, "Collector with id "+id+" has been deleted")
		}
	} else if offline == true {
		response, err := client.DeleteOfflineCollectors(&types.DeleteOfflineCollectorsOpts{
			AliveBeforeDays: optional.NewInt32(aliveBeforeDays),
		})
		if err != nil {
			cmdutils.OutputError(response, err)
		} else {
			cmdutils.Output(nil, response, err, "Offline collectors have been deleted")
		}
	}
}
