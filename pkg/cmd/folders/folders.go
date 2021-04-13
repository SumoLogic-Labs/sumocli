package folders

import (
	"github.com/spf13/cobra"
	cmdFoldersCreate "github.com/wizedkyle/sumocli/pkg/cmd/folders/create"
	cmdFoldersGet "github.com/wizedkyle/sumocli/pkg/cmd/folders/get"
	cmdFoldersPersonalFolder "github.com/wizedkyle/sumocli/pkg/cmd/folders/personal-folder"
	cmdFoldersUpdate "github.com/wizedkyle/sumocli/pkg/cmd/folders/update"
)

func NewCmdFolders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "folders <command>",
		Short: "Manage folders",
		Long:  "Commands that allow you to manage content in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdFoldersCreate.NewCmdCreate())
	cmd.AddCommand(cmdFoldersGet.NewCmdGet())
	cmd.AddCommand(cmdFoldersPersonalFolder.NewCmdPersonalFolder())
	cmd.AddCommand(cmdFoldersUpdate.NewCmdUpdate())
	return cmd
}
