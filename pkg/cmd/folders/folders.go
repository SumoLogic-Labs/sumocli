package folders

import (
	cmdFoldersCreate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/create"
	cmdFoldersGet "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/get"
	cmdFoldersAdminRecommendedFolder "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/get_admin_recommended_folder"
	cmdFoldersAdminRecommendedFolderResult "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/get_admin_recommended_folder_result"
	cmdFoldersAdminRecommendedFolderStatus "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/get_admin_recommended_folder_status"
	cmdFoldersGlobalFolder "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/get_global_folder"
	cmdFoldersGlobalFolderResult "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/get_global_folder_result"
	cmdFoldersGlobalFolderStatus "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/get_global_folder_status"
	cmdFoldersPersonalFolder "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/get_personal_folder"
	cmdFoldersUpdate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/folders/update"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdFolders(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "folders <command>",
		Short: "Manage folders",
		Long:  "Commands that allow you to manage content in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdFoldersAdminRecommendedFolder.NewCmdGetAdminRecommendedFolder(client))
	cmd.AddCommand(cmdFoldersAdminRecommendedFolderResult.NewCmdGetAdminRecommendedFolderResult(client))
	cmd.AddCommand(cmdFoldersAdminRecommendedFolderStatus.NewCmdGetAdminRecommendedFolderStatus(client))
	cmd.AddCommand(cmdFoldersCreate.NewCmdCreate(client))
	cmd.AddCommand(cmdFoldersGet.NewCmdGet(client))
	cmd.AddCommand(cmdFoldersGlobalFolder.NewCmdGetGlobalFolder(client))
	cmd.AddCommand(cmdFoldersGlobalFolderResult.NewCmdGetGlobalFolderResult(client))
	cmd.AddCommand(cmdFoldersGlobalFolderStatus.NewCmdGetGlobalFolderStatus(client))
	cmd.AddCommand(cmdFoldersPersonalFolder.NewCmdGetPersonalFolder(client))
	cmd.AddCommand(cmdFoldersUpdate.NewCmdUpdate(client))
	return cmd
}
