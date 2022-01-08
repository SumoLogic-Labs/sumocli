package get

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAppsGet(client *cip.APIClient) *cobra.Command {
	var uuid string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets the app with the given universally unique identifier (UUID).",
		Run: func(cmd *cobra.Command, args []string) {
			getApp(uuid, client)
		},
	}
	cmd.Flags().StringVar(&uuid, "uuid", "", "Specify the UUID of the app")
	cmd.MarkFlagRequired("uuid")
	return cmd
}

func getApp(uuid string, client *cip.APIClient) {
	data, response, err := client.GetApp(uuid)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
