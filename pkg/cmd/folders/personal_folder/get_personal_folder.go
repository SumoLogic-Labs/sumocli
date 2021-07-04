package get_personal_folder

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/config"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
)

func NewCmdGetPersonalFolder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-personal-folder",
		Short: "Get the personal folder of the current user.",
		Run: func(cmd *cobra.Command, args []string) {
			getPersonalFolder()
		},
	}
	return cmd
}

func getPersonalFolder() {
	client := config.GetSumoLogicSDKConfig()
	apiResponse, httpResponse, errorResponse := client.GetPersonalFolder()
	if errorResponse != nil {
		fmt.Println(errorResponse.Error())
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse)
	}
}
