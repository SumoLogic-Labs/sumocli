package create_field

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldManagementCreateField() *cobra.Command {
	var fieldName string

	cmd := &cobra.Command{
		Use:   "create-field",
		Short: "Adding a field will define it in the Fields schema allowing it to be assigned as metadata to your logs.",
		Run: func(cmd *cobra.Command, args []string) {
			createField(fieldName)
		},
	}
	cmd.Flags().StringVar(&fieldName, "fieldName", "", "Specify the name of the field")
	cmd.MarkFlagRequired("fieldName")
	return cmd
}

func createField(fieldName string) {
	var fieldsResponse api.Fields
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateField{
		FieldName: fieldName,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/fields"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
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
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(fieldsResponseJson))
	}
}
