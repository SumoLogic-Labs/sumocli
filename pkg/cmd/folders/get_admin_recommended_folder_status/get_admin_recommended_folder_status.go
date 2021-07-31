package get_admin_recommended_folder_status

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGetAdminRecommendedFolderStatus(client *cip.APIClient) *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "get-admin-recommended-folder-status",
		Short: "Get the status of an asynchronous Admin Recommended folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getAdminRecommendedFolderStatus(jobId, client)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id (returned from running sumocli admin-recommended-folder)")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getAdminRecommendedFolderStatus(jobId string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetAdminRecommendedFolderAsyncStatus(jobId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
