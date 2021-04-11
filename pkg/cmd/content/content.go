package content

import (
	"github.com/spf13/cobra"
	cmdContentDeletionStatus "github.com/wizedkyle/sumocli/pkg/cmd/content/deletion-status"
	cmdContentExportResult "github.com/wizedkyle/sumocli/pkg/cmd/content/export-result"
	cmdContentExportStatus "github.com/wizedkyle/sumocli/pkg/cmd/content/export-status"
	cmdContentGet "github.com/wizedkyle/sumocli/pkg/cmd/content/get"
	cmdContentGetPath "github.com/wizedkyle/sumocli/pkg/cmd/content/get-path"
	cmdContentImportStatus "github.com/wizedkyle/sumocli/pkg/cmd/content/import-status"
	cmdContentStartDeletion "github.com/wizedkyle/sumocli/pkg/cmd/content/start-deletion"
	cmdContentStartExport "github.com/wizedkyle/sumocli/pkg/cmd/content/start-export"
	cmdContentStartImport "github.com/wizedkyle/sumocli/pkg/cmd/content/start-import"
)

func NewCmdContent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content <command>",
		Short: "Manage content",
		Long:  "Commands that allow you to manage content in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdContentDeletionStatus.NewCmdDeletionStatus())
	cmd.AddCommand(cmdContentExportResult.NewCmdExportResult())
	cmd.AddCommand(cmdContentExportStatus.NewCmdExportStatus())
	cmd.AddCommand(cmdContentGet.NewCmdGet())
	cmd.AddCommand(cmdContentGetPath.NewCmdGetPath())
	cmd.AddCommand(cmdContentImportStatus.NewCmdImportStatus())
	cmd.AddCommand(cmdContentStartDeletion.NewCmdStartDeletion())
	cmd.AddCommand(cmdContentStartExport.NewCmdStartExport())
	cmd.AddCommand(cmdContentStartImport.NewCmdStartImport())
	return cmd
}
