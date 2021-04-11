package content

import (
	"github.com/spf13/cobra"
	cmdContentExportResult "github.com/wizedkyle/sumocli/pkg/cmd/content/export-result"
	cmdContentExportStatus "github.com/wizedkyle/sumocli/pkg/cmd/content/export-status"
	cmdContentGet "github.com/wizedkyle/sumocli/pkg/cmd/content/get"
	cmdContentGetPath "github.com/wizedkyle/sumocli/pkg/cmd/content/get-path"
	cmdContentStartExport "github.com/wizedkyle/sumocli/pkg/cmd/content/start-export"
)

func NewCmdContent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content <command>",
		Short: "Manage content",
		Long:  "Commands that allow you to manage content in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdContentExportResult.NewCmdExportResult())
	cmd.AddCommand(cmdContentExportStatus.NewCmdExportStatus())
	cmd.AddCommand(cmdContentGet.NewCmdGet())
	cmd.AddCommand(cmdContentGetPath.NewCmdGetPath())
	cmd.AddCommand(cmdContentStartExport.NewCmdStartExport())
	return cmd
}
