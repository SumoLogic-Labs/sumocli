package get_deletion_status

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
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
			authentication.ConfirmCredentialsSet(client)
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
	var options types.ContentOpts
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	data, response, err := client.GetAsyncDeleteStatus(id, jobId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
