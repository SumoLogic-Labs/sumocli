package get_admin_recommended_folder_result

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGetAdminRecommendedFolderResult(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "get-admin-recommended-folder-result",
		Short: "Get results from Admin Recommended job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getAdminRecommendedFolderResult(jobId, client, log)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id (returned from running sumocli admin-recommended-folder)")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getAdminRecommendedFolderResult(jobId string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.GetAdminRecommendedFolderAsyncResult(jobId)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get result of admin folder async job")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
