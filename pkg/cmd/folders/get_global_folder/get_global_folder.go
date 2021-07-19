package get_global_folder

import (
	"github.com/antihax/optional"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdGetGlobalFolder(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var isAdminMode bool

	cmd := &cobra.Command{
		Use: "get-global-folder",
		Short: "Schedule an asynchronous job to get global folder. " +
			"Global folder contains all content items that a user has permissions to view in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getGlobalFolder(isAdminMode, client, log)
		},
	}
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	return cmd
}

func getGlobalFolder(isAdminMode bool, client *cip.APIClient, log *zerolog.Logger) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	apiResponse, httpResponse, errorResponse := client.GetGlobalFolderAsync(&types.FolderManagementApiGetGlobalFolderAsyncOpts{
		IsAdminMode: optional.NewString(adminMode),
	})
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to crate global folder async job")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
