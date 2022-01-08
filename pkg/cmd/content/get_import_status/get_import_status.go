package get_import_status

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdGetImportStatus(client *cip.APIClient) *cobra.Command {
	var (
		folderId    string
		jobId       string
		isAdminMode bool
	)
	cmd := &cobra.Command{
		Use:   "import-status",
		Short: "Get the status of an asynchronous content import request for the given job identifier",
		Run: func(cmd *cobra.Command, args []string) {
			importStatus(folderId, jobId, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&folderId, "folderId", "", "Specify the id of the folder to import to")
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the import (returned from running sumocli content start-import)")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("folderId")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func importStatus(folderId string, jobId string, isAdminMode bool, client *cip.APIClient) {
	var options types.ContentOpts
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	data, response, err := client.GetAsyncImportStatus(folderId, jobId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
