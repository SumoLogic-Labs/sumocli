package get_personal_folder

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
