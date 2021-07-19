package delete_data

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdLookupTablesDeleteData() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete-data",
		Short: "Delete all data from a lookup table.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteLookupTableData(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table to delete data from")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteLookupTableData(id string) {
	var deleteDataResponse api.LookupTableRequestId
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/lookupTables/" + id + "/truncate"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &deleteDataResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	deleteLookupTableDataResponseJson, err := json.MarshalIndent(deleteDataResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(deleteLookupTableDataResponseJson))
	}
}
