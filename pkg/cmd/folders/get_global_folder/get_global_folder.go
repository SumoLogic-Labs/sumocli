package get_global_folder

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdGetGlobalFolder(client *cip.APIClient) *cobra.Command {
	var isAdminMode bool

	cmd := &cobra.Command{
		Use: "get-global-folder",
		Short: "Schedule an asynchronous job to get global folder. " +
			"Global folder contains all content items that a user has permissions to view in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getGlobalFolder(isAdminMode, client)
		},
	}
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	return cmd
}

func getGlobalFolder(isAdminMode bool, client *cip.APIClient) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	apiResponse, httpResponse, errorResponse := client.GetGlobalFolderAsync(&types.FolderOpts{
		IsAdminMode: optional.NewString(adminMode),
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
