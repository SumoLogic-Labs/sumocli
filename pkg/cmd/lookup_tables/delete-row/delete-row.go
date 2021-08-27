package delete_row

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdLookupTablesDeleteRow(client *cip.APIClient) *cobra.Command {
	var (
		id           string
		columnNames  []string
		columnValues []string
	)
	cmd := &cobra.Command{
		Use:   "delete-row",
		Short: "Delete a row from lookup table by providing the primary key. The complete set of primary key fields of the lookup table should be provided.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteLookupTableRow(id, columnNames, columnValues, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table")
	cmd.Flags().StringSliceVar(&columnNames, "columnNames", []string{}, "List of primary key column names "+
		"(they need to be comma separated e.g. field1,field2,field3)")
	cmd.Flags().StringSliceVar(&columnValues, "columnValues", []string{}, "List of values "+
		"(they need to be comma separated e.g. value1,value2,value3).")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("columnNames")
	cmd.MarkFlagRequired("columnValues")
	return cmd
}

func deleteLookupTableRow(id string, columnNames []string, columnValues []string, client *cip.APIClient) {
	httpResponse, errorResponse := client.DeleteTableRow(types.RowDeleteDefinition{
		PrimaryKey: cmdutils.GenerateLookupTableColumns(columnNames, columnValues),
	},
		id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Row deleted successfully.")
	}
}
