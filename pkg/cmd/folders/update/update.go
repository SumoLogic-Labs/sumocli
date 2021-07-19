package update

import (
	"github.com/antihax/optional"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdUpdate(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
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
			update(name, description, id, isAdminMode, client, log)
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

func update(name string, description string, id string, isAdminMode bool, client *cip.APIClient, log *zerolog.Logger) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	apiResponse, httpResponse, errorResponse := client.UpdateFolder(types.UpdateFolderRequest{
		Name:        name,
		Description: description,
	}, id, &types.FolderManagementApiUpdateFolderOpts{
		IsAdminMode: optional.NewString(adminMode),
	})
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to update folder")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
