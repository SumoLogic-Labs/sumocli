package get_global_folder_status

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdGetGlobalFolderStatus(client *cip.APIClient) *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "get-global-folder-status",
		Short: "Get the status of an asynchronous global folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getGlobalFolderStatus(jobId, client)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the global folder (returned from running sumocli folders global-folder")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getGlobalFolderStatus(jobId string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetGlobalFolderAsyncStatus(jobId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
