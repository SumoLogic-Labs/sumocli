package get_admin_recommended_folder_status

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGetAdminRecommendedFolderStatus(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "get-admin-recommended-folder-status",
		Short: "Get the status of an asynchronous Admin Recommended folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getAdminRecommendedFolderStatus(jobId, client, log)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id (returned from running sumocli admin-recommended-folder)")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getAdminRecommendedFolderStatus(jobId string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.GetAdminRecommendedFolderAsyncStatus(jobId)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get admin folder async job status")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
