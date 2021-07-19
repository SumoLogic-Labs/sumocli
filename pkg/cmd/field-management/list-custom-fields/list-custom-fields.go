package list_custom_fields

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldManagementListCustomFields() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-custom-fields",
		Short: "Request a list of all the custom fields configured in your account.",
		Run: func(cmd *cobra.Command, args []string) {
			listCustomFields()
		},
	}
	return cmd
}

func listCustomFields() {
	var fieldsResponse api.GetFields
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/fields"
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
