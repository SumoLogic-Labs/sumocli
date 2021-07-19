package get_admin_recommended_folder

import (
	"github.com/antihax/optional"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdGetAdminRecommendedFolder(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var isAdminMode bool

	cmd := &cobra.Command{
		Use:   "get-admin-recommended-folder",
		Short: "Schedule an asynchronous job to get the top-level Admin Recommended content items.",
		Run: func(cmd *cobra.Command, args []string) {
			getAdminRecommendedFolder(isAdminMode, client, log)
		},
	}
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	return cmd
}

func getAdminRecommendedFolder(isAdminMode bool, client *cip.APIClient, log *zerolog.Logger) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	apiResponse, httpResponse, errorResponse := client.GetAdminRecommendedFolderAsync(&types.FolderManagementApiGetAdminRecommendedFolderAsyncOpts{
		IsAdminMode: optional.NewString(adminMode),
	})
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to create admin folder async job")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
