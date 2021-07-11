package folders

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	cmdFoldersCreate "github.com/wizedkyle/sumocli/pkg/cmd/folders/create"
	cmdFoldersGet "github.com/wizedkyle/sumocli/pkg/cmd/folders/get"
	cmdFoldersAdminRecommendedFolder "github.com/wizedkyle/sumocli/pkg/cmd/folders/get_admin_recommended_folder"
	cmdFoldersAdminRecommendedFolderResult "github.com/wizedkyle/sumocli/pkg/cmd/folders/get_admin_recommended_folder_result"
	cmdFoldersAdminRecommendedFolderStatus "github.com/wizedkyle/sumocli/pkg/cmd/folders/get_admin_recommended_folder_status"
	cmdFoldersGlobalFolder "github.com/wizedkyle/sumocli/pkg/cmd/folders/get_global_folder"
	cmdFoldersGlobalFolderResult "github.com/wizedkyle/sumocli/pkg/cmd/folders/get_global_folder_result"
	cmdFoldersGlobalFolderStatus "github.com/wizedkyle/sumocli/pkg/cmd/folders/get_global_folder_status"
	cmdFoldersPersonalFolder "github.com/wizedkyle/sumocli/pkg/cmd/folders/get_personal_folder"
	cmdFoldersUpdate "github.com/wizedkyle/sumocli/pkg/cmd/folders/update"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdFolders(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "folders <command>",
		Short: "Manage folders",
		Long:  "Commands that allow you to manage content in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdFoldersAdminRecommendedFolder.NewCmdGetAdminRecommendedFolder(client, log))
	cmd.AddCommand(cmdFoldersAdminRecommendedFolderResult.NewCmdGetAdminRecommendedFolderResult(client, log))
	cmd.AddCommand(cmdFoldersAdminRecommendedFolderStatus.NewCmdGetAdminRecommendedFolderStatus(client, log))
	cmd.AddCommand(cmdFoldersCreate.NewCmdCreate(client, log))
	cmd.AddCommand(cmdFoldersGet.NewCmdGet(client, log))
	cmd.AddCommand(cmdFoldersGlobalFolder.NewCmdGetGlobalFolder(client, log))
	cmd.AddCommand(cmdFoldersGlobalFolderResult.NewCmdGetGlobalFolderResult(client, log))
	cmd.AddCommand(cmdFoldersGlobalFolderStatus.NewCmdGetGlobalFolderStatus(client, log))
	cmd.AddCommand(cmdFoldersPersonalFolder.NewCmdGetPersonalFolder(client, log))
	cmd.AddCommand(cmdFoldersUpdate.NewCmdUpdate(client, log))
	return cmd
}
