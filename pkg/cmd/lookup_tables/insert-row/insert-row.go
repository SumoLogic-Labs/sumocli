package insert_row

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdLookupTablesInsertRow(client *cip.APIClient) *cobra.Command {
	var (
		id           string
		columnNames  []string
		columnValues []string
	)
	cmd := &cobra.Command{
		Use:   "insert-row",
		Short: "Insert or update a row of a lookup table with the given identifier.",
		Long: "A new row is inserted if the primary key does not exist already, otherwise the existing row with the specified primary key is updated. All the fields of the lookup table are required and will be updated to the given values. " +
			"In case a field is not specified then it will be assumed to be set to null.",
		Run: func(cmd *cobra.Command, args []string) {
			insertLookupTableRow(id, columnNames, columnValues, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table")
	cmd.Flags().StringSliceVar(&columnNames, "columnNames", []string{}, "List of column names "+
		"(they need to be comma separated e.g. field1,field2,field3)")
	cmd.Flags().StringSliceVar(&columnValues, "columnValues", []string{}, "List of values "+
		"(they need to be comma separated e.g. value1,value2,value3).")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("columnNames")
	cmd.MarkFlagRequired("columnValues")
	return cmd
}

func insertLookupTableRow(id string, columnNames []string, columnValues []string, client *cip.APIClient) {
	httpResponse, errorResponse := client.UpdateTableRow(types.RowUpdateDefinition{
		Row: cmdutils.GenerateLookupTableColumns(columnNames, columnValues),
	},
		id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Row updated successfully.")
	}
}
