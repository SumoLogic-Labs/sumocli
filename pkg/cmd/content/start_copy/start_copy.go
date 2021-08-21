package start_copy

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdStartCopy(client *cip.APIClient) *cobra.Command {
	var (
		id                string
		destinationFolder string
		isAdminMode       bool
	)
	cmd := &cobra.Command{
		Use:   "start-copy",
		Short: "Start an asynchronous content copy job with the given identifier to the destination folder. If the content item is a folder, everything under the folder is copied recursively.",
		Run: func(cmd *cobra.Command, args []string) {
			startCopy(id, destinationFolder, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content you want to copy")
	cmd.Flags().StringVar(&destinationFolder, "destinationFolder", "", "Specify the id of the destination folder")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "et to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("destinationFolder")
	return cmd
}

func startCopy(id string, destinationFolder string, isAdminMode bool, client *cip.APIClient) {
	var options types.ContentOpts
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	apiResponse, httpResponse, errorResponse := client.BeginAsyncCopy(id, destinationFolder, &options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
