package update

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
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
	apiResponse, httpResponse, errorResponse := client.UpdateFolder(types.UpdateFolderRequest{
		Name:        name,
		Description: description,
	}, id, &types.FolderOpts{
		IsAdminMode: optional.NewString(adminMode),
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
