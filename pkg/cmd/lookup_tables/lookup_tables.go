package lookup_tables

import (
	cmdLookupTablesCreate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables/create"
	cmdLookupTablesDelete "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables/delete"
	cmdLookupTablesDeleteData "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables/delete-data"
	cmdLookupTablesDeleteRow "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables/delete-row"
	cmdLookupTablesGet "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables/get"
	cmdLookupTablesInsertRow "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables/insert-row"
	cmdLookupTablesJobStatus "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables/job-status"
	cmdLookupTablesUpdate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables/update"
	cmdLookupTablesUpload "github.com/SumoLogic-Labs/sumocli/pkg/cmd/lookup_tables/upload"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdLookupTables(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lookup-tables",
		Short: "Manage lookup tables",
		Long:  "Commands that allow you to manage Lookup Tables in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdLookupTablesCreate.NewCmdLookupTablesCreate(client))
	cmd.AddCommand(cmdLookupTablesDelete.NewCmdLookupTablesDelete(client))
	cmd.AddCommand(cmdLookupTablesDeleteData.NewCmdLookupTablesDeleteData(client))
	cmd.AddCommand(cmdLookupTablesDeleteRow.NewCmdLookupTablesDeleteRow(client))
	cmd.AddCommand(cmdLookupTablesUpdate.NewCmdLookupTablesEdit(client))
	cmd.AddCommand(cmdLookupTablesGet.NewCmdLookupTablesGet(client))
	cmd.AddCommand(cmdLookupTablesInsertRow.NewCmdLookupTablesInsertRow(client))
	cmd.AddCommand(cmdLookupTablesJobStatus.NewCmdLookupTableJobStatus(client))
	cmd.AddCommand(cmdLookupTablesUpload.NewCmdLookupTablesUpload(client))
	return cmd
}
