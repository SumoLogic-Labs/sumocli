package lookup_tables

import (
	"github.com/spf13/cobra"
	cmdLookupTablesGet "github.com/wizedkyle/sumocli/pkg/cmd/lookup-tables/get"
)

func NewCmdLookupTables() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lookup-tables",
		Short: "Manage lookup tables",
		Long:  "Commands that allow you to manage Lookup Tables in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdLookupTablesGet.NewCmdLookupTablesGet())
	return cmd
}
