package content

import (
	"github.com/spf13/cobra"
	cmdContentGet "github.com/wizedkyle/sumocli/pkg/cmd/content/get"
	cmdContentCopyStatus "github.com/wizedkyle/sumocli/pkg/cmd/content/get_copy_status"
	cmdContentDeletionStatus "github.com/wizedkyle/sumocli/pkg/cmd/content/get_deletion_status"
	cmdContentExportResult "github.com/wizedkyle/sumocli/pkg/cmd/content/get_export_result"
	cmdContentExportStatus "github.com/wizedkyle/sumocli/pkg/cmd/content/get_export_status"
	cmdContentImportStatus "github.com/wizedkyle/sumocli/pkg/cmd/content/import-status"
	cmdContentMove "github.com/wizedkyle/sumocli/pkg/cmd/content/move"
	cmdContentStartCopy "github.com/wizedkyle/sumocli/pkg/cmd/content/start-copy"
	cmdContentStartDeletion "github.com/wizedkyle/sumocli/pkg/cmd/content/start-deletion"
	cmdContentStartImport "github.com/wizedkyle/sumocli/pkg/cmd/content/start-import"
	cmdContentStartExport "github.com/wizedkyle/sumocli/pkg/cmd/content/start_export"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdContent(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content <command>",
		Short: "Manage content",
		Long:  "Commands that allow you to manage content in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdContentCopyStatus.NewCmdGetCopyStatus(client))
	cmd.AddCommand(cmdContentDeletionStatus.NewCmdGetDeletionStatus(client))
	cmd.AddCommand(cmdContentExportResult.NewCmdGetExportResult(client))
	cmd.AddCommand(cmdContentExportStatus.NewCmdExportStatus(client))
	cmd.AddCommand(cmdContentGet.NewCmdGet(client))
	cmd.AddCommand(cmdContentImportStatus.NewCmdImportStatus())
	cmd.AddCommand(cmdContentMove.NewCmdMove())
	cmd.AddCommand(cmdContentStartDeletion.NewCmdStartDeletion())
	cmd.AddCommand(cmdContentStartCopy.NewCmdStartCopy())
	cmd.AddCommand(cmdContentStartExport.NewCmdStartExport(client))
	cmd.AddCommand(cmdContentStartImport.NewCmdStartImport())
	return cmd
}
