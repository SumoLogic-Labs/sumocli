package update

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdUpdate(client *cip.APIClient) *cobra.Command {
	var (
		name        string
		description string
		id          string
		isAdminMode bool
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing folder with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			update(name, description, id, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the folder")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the folder")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the folder to update")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("id")
	return cmd
}

func update(name string, description string, id string, isAdminMode bool, client *cip.APIClient) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	data, response, err := client.UpdateFolder(types.UpdateFolderRequest{
		Name:        name,
		Description: description,
	}, id, &types.FolderOpts{
		IsAdminMode: optional.NewString(adminMode),
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
