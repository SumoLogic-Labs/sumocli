package get_available_builds

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
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
	apiResponse, httpResponse, errorResponse := client.GetAvailableBuilds()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
