package list

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
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
	data, response, err := client.ListApps()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
