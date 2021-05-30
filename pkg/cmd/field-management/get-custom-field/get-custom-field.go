package get_custom_field

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldManagementGetCustomField() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get-custom-field",
		Short: "Get the details of a custom field.",
		Run: func(cmd *cobra.Command, args []string) {
			getCustomField(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the custom field")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getCustomField(id string) {
	var fieldsResponse api.Fields
	log := logging.GetConsoleLogger()
	requestUrl := "v1/fields/" + id
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
