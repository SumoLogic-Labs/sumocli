package folders

import (
	"github.com/spf13/cobra"
	cmdFoldersAdminRecommendedFolder "github.com/wizedkyle/sumocli/pkg/cmd/folders/admin-recommended-folder"
	cmdFoldersAdminRecommendedFolderResult "github.com/wizedkyle/sumocli/pkg/cmd/folders/admin-recommended-folder-result"
	cmdFoldersAdminRecommendedFolderStatus "github.com/wizedkyle/sumocli/pkg/cmd/folders/admin-recommended-folder-status"
	cmdFoldersCreate "github.com/wizedkyle/sumocli/pkg/cmd/folders/create"
	cmdFoldersGet "github.com/wizedkyle/sumocli/pkg/cmd/folders/get"
	cmdFoldersPersonalFolder "github.com/wizedkyle/sumocli/pkg/cmd/folders/get_personal_folder"
	cmdFoldersGlobalFolder "github.com/wizedkyle/sumocli/pkg/cmd/folders/global-folder"
	cmdFoldersGlobalFolderResult "github.com/wizedkyle/sumocli/pkg/cmd/folders/global-folder-result"
	cmdFoldersGlobalFolderStatus "github.com/wizedkyle/sumocli/pkg/cmd/folders/global-folder-status"
	cmdFoldersUpdate "github.com/wizedkyle/sumocli/pkg/cmd/folders/update"
)

func NewCmdFolders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "folders <command>",
		Short: "Manage folders",
		Long:  "Commands that allow you to manage content in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdFoldersAdminRecommendedFolder.NewCmdAdminRecommendedFolder())
	cmd.AddCommand(cmdFoldersAdminRecommendedFolderResult.NewCmdAdminRecommendedFolderResult())
	cmd.AddCommand(cmdFoldersAdminRecommendedFolderStatus.NewCmdAdminRecommendedFolderStatus())
	cmd.AddCommand(cmdFoldersCreate.NewCmdCreate())
	cmd.AddCommand(cmdFoldersGet.NewCmdGet())
	cmd.AddCommand(cmdFoldersGlobalFolder.NewCmdGlobalFolder())
	cmd.AddCommand(cmdFoldersGlobalFolderResult.NewCmdGlobalFolderResult())
	cmd.AddCommand(cmdFoldersGlobalFolderStatus.NewCmdGlobalFolderStatus())
	cmd.AddCommand(cmdFoldersPersonalFolder.NewCmdGetPersonalFolder())
	cmd.AddCommand(cmdFoldersUpdate.NewCmdUpdate())
	return cmd
}
