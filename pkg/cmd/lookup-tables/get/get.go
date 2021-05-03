package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdLookupTablesGet() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic lookup table based on the given identifier",
		Run: func(cmd *cobra.Command, args []string) {
			getLookupTable(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table you want to retrieve")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getLookupTable(id string) {
	var lookupTableResponse api.LookupTableResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v1/lookupTables/" + id
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &lookupTableResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	lookupTableResponseJson, err := json.MarshalIndent(lookupTableResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(lookupTableResponseJson))
	}
}
