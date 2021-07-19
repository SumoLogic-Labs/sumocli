package empty

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdLookupTableEmpty() *cobra.Command {
	var (
		id         string
		jsonFormat bool
	)

	cmd := &cobra.Command{
		Use:   "empty",
		Short: "Delete all data from a lookup table.",
		Run: func(cmd *cobra.Command, args []string) {
			emptyLookupTable(id, jsonFormat)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table to empty")
	cmd.Flags().BoolVar(&jsonFormat, "jsonFormat", false, "Set to true if you want the output to be formatted JSON")
	cmd.MarkFlagRequired("id")
	return cmd
}

func emptyLookupTable(id string, jsonFormat bool) {
	var emptyLookupTableRequest api.LookupTableRequestId
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/lookupTables/" + id + "/truncate"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &emptyLookupTableRequest)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		if jsonFormat == true {
			emptyLookupTableRequestJson, err := json.MarshalIndent(emptyLookupTableRequest, "", "    ")
			if err != nil {
				log.Error().Err(err).Msg("failed to marshal response")
			}
			fmt.Print(string(emptyLookupTableRequestJson))
		} else {
			emptyLookupTableRequestJson, err := json.Marshal(emptyLookupTableRequest)
			if err != nil {
				log.Error().Err(err).Msg("failed to marshal response")
			}
			fmt.Println(string(emptyLookupTableRequestJson))
		}
	}
}
