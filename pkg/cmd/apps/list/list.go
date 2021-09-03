package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAppsList(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all available apps from the App Catalog.",
		Run: func(cmd *cobra.Command, args []string) {
			listAvailableApps(client)
		},
	}
	return cmd
}

func listAvailableApps(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.ListApps()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
