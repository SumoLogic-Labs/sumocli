package move

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
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
			authentication.ConfirmCredentialsSet(client)
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
	response, err := client.MoveItem(destinationFolderId, id, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Content was moved successfully.")
	}
}
