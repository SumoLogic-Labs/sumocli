package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
	apiResponse, httpResponse, errorResponse := client.GetApp(uuid)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
