package get_export_status

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdExportStatus(client *cip.APIClient) *cobra.Command {
	var (
		contentId   string
		jobId       string
		isAdminMode bool
	)
	cmd := &cobra.Command{
		Use:   "get-export-status",
		Short: "Get the status of an asynchronous content export request for the given job identifier",
		Run: func(cmd *cobra.Command, args []string) {
			exportStatus(contentId, jobId, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&contentId, "contentId", "", "Specify the id of the content item to export")
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the export (returned from running sumocli content start-export)")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("contentId")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func exportStatus(contentId string, jobId string, isAdminMode bool, client *cip.APIClient) {
	var options types.ContentOpts
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	apiResponse, httpResponse, errorResponse := client.GetAsyncExportStatus(contentId, jobId, &options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
