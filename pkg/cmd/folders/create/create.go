package create

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdCreate(client *cip.APIClient) *cobra.Command {
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
			create(name, description, parentId, isAdminMode, client)
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

func create(name string, description string, parentId string, isAdminMode bool, client *cip.APIClient) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	data, response, err := client.CreateFolder(types.FolderDefinition{
		Name:        name,
		Description: description,
		ParentId:    parentId,
	},
		&types.FolderOpts{
			IsAdminMode: optional.NewString(adminMode),
		})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
