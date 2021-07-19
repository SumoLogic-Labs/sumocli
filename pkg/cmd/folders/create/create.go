package create

import (
	"github.com/antihax/optional"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdCreate(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var (
		name        string
		description string
		parentId    string
		isAdminMode bool
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new folder under the given parent folder.",
		Run: func(cmd *cobra.Command, args []string) {
			create(name, description, parentId, isAdminMode, client, log)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the folder")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the folder")
	cmd.Flags().StringVar(&parentId, "parentId", "", "Specify the parent folder Id")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("parentId")
	return cmd
}

func create(name string, description string, parentId string, isAdminMode bool, client *cip.APIClient, log *zerolog.Logger) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	apiResponse, httpResponse, errorResponse := client.CreateFolder(types.FolderDefinition{
		Name:        name,
		Description: description,
		ParentId:    parentId,
	},
		&types.FolderManagementApiCreateFolderOpts{
			IsAdminMode: optional.NewString(adminMode),
		})
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to create folder")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
