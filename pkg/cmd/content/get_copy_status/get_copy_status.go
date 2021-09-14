package get_copy_status

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
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
	data, response, err := client.AsyncCopyStatus(id, jobId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
