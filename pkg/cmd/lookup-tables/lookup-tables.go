package lookup_tables

import (
	"github.com/spf13/cobra"
	cmdLookupTablesCreate "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/create"
	cmdLookupTablesDelete "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/delete"
	cmdLookupTablesDeleteData "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/delete-data"
	cmdLookupTablesDeleteRow "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/delete-row"
	cmdLookupTablesEdit "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/edit"
	cmdLookupTablesEmpty "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/empty"
	cmdLookupTablesGet "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/get"
	cmdLookupTablesInsertRow "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/insert-row"
	cmdLookupTablesJobStatus "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/job-status"
)

func NewCmdLookupTables() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lookup-tables",
		Short: "Manage lookup tables",
		Long:  "Commands that allow you to manage Lookup Tables in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdLookupTablesCreate.NewCmdLookupTablesCreate())
	cmd.AddCommand(cmdLookupTablesDelete.NewCmdLookupTablesDelete())
	cmd.AddCommand(cmdLookupTablesDeleteData.NewCmdLookupTablesDeleteData())
	cmd.AddCommand(cmdLookupTablesDeleteRow.NewCmdLookupTablesDeleteRow())
	cmd.AddCommand(cmdLookupTablesEdit.NewCmdLookupTablesEdit())
	cmd.AddCommand(cmdLookupTablesEmpty.NewCmdLookupTableEmpty())
	cmd.AddCommand(cmdLookupTablesGet.NewCmdLookupTablesGet())
	cmd.AddCommand(cmdLookupTablesInsertRow.NewCmdLookupTablesInsertRow())
	cmd.AddCommand(cmdLookupTablesJobStatus.NewCmdLookupTableJobStatus())
	return cmd
}
