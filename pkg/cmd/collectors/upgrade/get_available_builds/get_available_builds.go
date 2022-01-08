package get_available_builds

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdGetBuilds(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-available-builds",
		Short: "Gets available Sumo Logic collector builds",
		Run: func(cmd *cobra.Command, args []string) {
			getBuilds(client)
		},
	}
	return cmd
}

func getBuilds(client *cip.APIClient) {
	data, response, err := client.GetAvailableBuilds()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
