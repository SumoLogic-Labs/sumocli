package get_deletion_status

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdGetDeletionStatus(client *cip.APIClient) *cobra.Command {
	var (
		id          string
		jobId       string
		isAdminMode bool
	)
	cmd := &cobra.Command{
		Use:   "get-deletion-status",
		Short: "Get the status of an asynchronous content deletion job request for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			deletionStatus(id, jobId, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content to delete")
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the deletion (returned from running sumocli content start-deletion)")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("contentId")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func deletionStatus(id string, jobId string, isAdminMode bool, client *cip.APIClient) {
	var options types.ContentManagementApiGetAsyncDeleteStatusOpts
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	apiResponse, httpResponse, errorResponse := client.GetAsyncDeleteStatus(id, jobId, &options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
