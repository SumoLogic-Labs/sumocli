package job_status

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdLookupTableJobStatus(client *cip.APIClient) *cobra.Command {
	var jobId string
	cmd := &cobra.Command{
		Use:   "job-status",
		Short: "Retrieve the status of a previously made request using sumocli lookup-tables upload or sumocli lookup-tables empty",
		Run: func(cmd *cobra.Command, args []string) {
			getLookupTableJobStatus(jobId, client)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the Job ID to get the status for (returned from running sumocli lookup-tables upload or sumocli lookup-tables empty")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getLookupTableJobStatus(jobId string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.RequestJobStatus(jobId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
