package get_copy_status

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdGetCopyStatus(client *cip.APIClient) *cobra.Command {
	var (
		id          string
		jobId       string
		isAdminMode bool
	)
	cmd := &cobra.Command{
		Use:   "get-copy-status",
		Short: "Get the status of the copy request with the given job identifier. On success, field statusMessage will contain identifier of the newly copied content.",
		Run: func(cmd *cobra.Command, args []string) {
			copyStatus(id, jobId, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content that was copied")
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the copy (returned from running sumocli content start-copy)")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func copyStatus(id string, jobId string, isAdminMode bool, client *cip.APIClient) {
	var options types.ContentOpts
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	apiResponse, httpResponse, errorResponse := client.AsyncCopyStatus(id, jobId, &options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
