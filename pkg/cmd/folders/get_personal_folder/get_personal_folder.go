package get_personal_folder

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGetPersonalFolder(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-personal-folder",
		Short: "Get the personal folder of the current user.",
		Run: func(cmd *cobra.Command, args []string) {
			getPersonalFolder(client, log)
		},
	}
	return cmd
}

func getPersonalFolder(client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.GetPersonalFolder()
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get personal folder")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
