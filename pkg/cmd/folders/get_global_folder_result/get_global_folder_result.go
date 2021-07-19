package get_global_folder_result

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGetGlobalFolderResult(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "get-global-folder-result",
		Short: "Get results from global folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getGlobalFolderResult(jobId, client, log)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the global folder (returned from running sumocli folders global-folder")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getGlobalFolderResult(jobId string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.GetGlobalFolderAsyncResult(jobId)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get global folder async job result")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
