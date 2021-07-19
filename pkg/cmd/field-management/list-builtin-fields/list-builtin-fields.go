package list_builtin_fields

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldManagementListBuiltinFields() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-builtin-fields",
		Short: "Built-in fields are created automatically by Sumo Logic for standard configuration purposes. " +
			"They include _sourceHost and _sourceCategory. Built-in fields can't be deleted or disabled.",
		Run: func(cmd *cobra.Command, args []string) {
			listBuiltinFields()
		},
	}
	return cmd
}

func listBuiltinFields() {
	var fieldsResponse api.GetFields
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/fields/builtin"
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

	err = json.Unmarshal(responseBody, &fieldsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	fieldsResponseJson, err := json.MarshalIndent(fieldsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(fieldsResponseJson))
	}
}
