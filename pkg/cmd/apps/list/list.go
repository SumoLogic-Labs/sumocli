package list

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAppsList(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all available apps from the App Catalog.",
		Run: func(cmd *cobra.Command, args []string) {
			listAvailableApps(client, log)
		},
	}
	return cmd
}

func listAvailableApps(client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.ListApps()
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to list apps")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
