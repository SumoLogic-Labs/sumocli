package move

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdMove(client *cip.APIClient) *cobra.Command {
	var (
		id                  string
		destinationFolderId string
		isAdminMode         bool
	)
	cmd := &cobra.Command{
		Use:   "move",
		Short: "Moves an item from its current location to another folder.",
		Run: func(cmd *cobra.Command, args []string) {
			move(id, destinationFolderId, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content to move")
	cmd.Flags().StringVar(&destinationFolderId, "destinationFolderId", "", "Specify the destination folder to move the content to")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("destinationFolderId")
	return cmd
}

func move(id string, destinationFolderId string, isAdminMode bool, client *cip.APIClient) {
	var options types.ContentOpts
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	httpResponse, errorResponse := client.MoveItem(destinationFolderId, id, &options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Content was moved successfully.")
	}
}
