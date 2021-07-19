package insert_row

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"strings"
)

func NewCmdLookupTablesInsertRow() *cobra.Command {
	var (
		id           string
		columnNames  string
		columnValues string
	)

	cmd := &cobra.Command{
		Use:   "insert-row",
		Short: "Insert or update a row of a lookup table with the given identifier.",
		Long: " new row is inserted if the primary key does not exist already, otherwise the existing row with the specified primary key is updated. All the fields of the lookup table are required and will be updated to the given values. " +
			"In case a field is not specified then it will be assumed to be set to null.",
		Run: func(cmd *cobra.Command, args []string) {
			insertLookupTableRow(id, columnNames, columnValues)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table")
	cmd.Flags().StringVar(&columnNames, "columnNames", "", "List of column names "+
		"(they need to be comma separated e.g. field1,field2,field3)")
	cmd.Flags().StringVar(&columnValues, "columnValues", "", "List of values "+
		"(they need to be comma separated e.g. value1,value2,value3).")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("columnNames")
	cmd.MarkFlagRequired("columnValues")
	return cmd
}

func insertLookupTableRow(id string, columnNames string, columnValues string) {
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.LookupTableRowRequest{}
	columnNameSlice := strings.Split(columnNames, ",")
	columnValueSlice := strings.Split(columnValues, ",")
	for i := range columnNameSlice {
		columnAddition := api.LookupTableRow{
			ColumnName:  columnNameSlice[i],
			ColumnValue: columnValueSlice[i],
		}
		requestBodySchema.Row = append(requestBodySchema.Row, columnAddition)
		i++
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "/v1/lookupTables/" + id + "/row"
	client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	if response.StatusCode == 204 {
		fmt.Println("Row updated successfully")
	} else {
		factory.HttpError(response.StatusCode, responseBody, log)
	}
}
