package get_admin_recommended_folder_status

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
	data, response, err := client.GetAdminRecommendedFolderAsyncStatus(jobId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
