package builds

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGetBuilds(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "builds",
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
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
