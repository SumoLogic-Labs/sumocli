package delete_row

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

func NewCmdLookupTablesDeleteRow() *cobra.Command {
	var (
		id           string
		columnNames  string
		columnValues string
	)

	cmd := &cobra.Command{
		Use:   "delete-row",
		Short: "Delete a row from lookup table by giving primary key. The complete set of primary key fields of the lookup table should be provided.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteLookupTableRow(id, columnNames, columnValues)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table")
	cmd.Flags().StringVar(&columnNames, "columnNames", "", "List of primary key column names "+
		"(they need to be comma separated e.g. field1,field2,field3)")
	cmd.Flags().StringVar(&columnValues, "columnValues", "", "List of values "+
		"(they need to be comma separated e.g. value1,value2,value3).")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("columnNames")
	cmd.MarkFlagRequired("columnValues")
	return cmd
}

func deleteLookupTableRow(id string, columnNames string, columnValues string) {
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.LookupTableRemoveRowRequest{}
	columnNameSlice := strings.Split(columnNames, ",")
	columnValueSlice := strings.Split(columnValues, ",")
	for i := range columnNameSlice {
		columnAddition := api.LookupTableRow{
			ColumnName:  columnNameSlice[i],
			ColumnValue: columnValueSlice[i],
		}
		requestBodySchema.PrimaryKey = append(requestBodySchema.PrimaryKey, columnAddition)
		i++
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "/v1/lookupTables/" + id + "/deleteTableRow"
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
		fmt.Println("Row deleted successfully")
	} else {
		factory.HttpError(response.StatusCode, responseBody, log)
	}
}
