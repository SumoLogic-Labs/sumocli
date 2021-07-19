package list_dropped_fields

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldManagementListDroppedFields() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-dropped-fields",
		Short: "Dropped fields are fields sent to Sumo Logic, but are ignored since they are not defined in your Fields schema. " +
			"In order to save these values a field must both exist and be enabled.",
		Run: func(cmd *cobra.Command, args []string) {
			listDroppedFields()
		},
	}
	return cmd
}

func listDroppedFields() {
	var fieldsResponse api.GetDroppedFields
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/fields/dropped"
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
