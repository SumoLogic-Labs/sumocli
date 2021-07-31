package get_global_folder_result

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGetGlobalFolderResult(client *cip.APIClient) *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "get-global-folder-result",
		Short: "Get results from global folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getGlobalFolderResult(jobId, client)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the global folder (returned from running sumocli folders global-folder")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getGlobalFolderResult(jobId string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetGlobalFolderAsyncResult(jobId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
