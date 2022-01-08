package install_status

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAppsInstallStatus(client *cip.APIClient) *cobra.Command {
	var jobId string
	cmd := &cobra.Command{
		Use:   "install-status",
		Short: "Get the status of an asynchronous app install request for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getAppInstallStatus(jobId, client)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify a jobId (it can be retrieved by running sumocli apps install)")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getAppInstallStatus(jobId string, client *cip.APIClient) {
	data, response, err := client.GetAsyncInstallStatus(jobId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
