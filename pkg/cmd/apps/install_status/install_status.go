package install_status

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAppsInstallStatus(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var jobId string
	cmd := &cobra.Command{
		Use:   "install-status",
		Short: "Get the status of an asynchronous app install request for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getAppInstallStatus(jobId, client, log)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify a jobId (it can be retrieved by running sumocli apps install)")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getAppInstallStatus(jobId string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.GetAsyncInstallStatus(jobId)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get install status")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
