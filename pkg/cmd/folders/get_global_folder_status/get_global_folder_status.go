package get_global_folder_status

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGetGlobalFolderStatus(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "get-global-folder-status",
		Short: "Get the status of an asynchronous global folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getGlobalFolderStatus(jobId, client, log)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the global folder (returned from running sumocli folders global-folder")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getGlobalFolderStatus(jobId string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.GetGlobalFolderAsyncStatus(jobId)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get global folder job status")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
