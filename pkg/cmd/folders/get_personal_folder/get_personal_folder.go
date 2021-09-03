package get_personal_folder

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdGetPersonalFolder(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-personal-folder",
		Short: "Get the personal folder of the current user.",
		Run: func(cmd *cobra.Command, args []string) {
			getPersonalFolder(client)
		},
	}
	return cmd
}

func getPersonalFolder(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetPersonalFolder()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
