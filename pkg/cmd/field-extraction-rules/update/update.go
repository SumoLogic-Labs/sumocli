package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldExtractionRulesUpdate() *cobra.Command {
	var (
		id              string
		name            string
		scope           string
		parseExpression string
		enabled         bool
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing field extraction rule.",
		Run: func(cmd *cobra.Command, args []string) {
			updateFieldExtractionRule(id, name, scope, parseExpression, enabled)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field extraction rule")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the field extraction rule")
	cmd.Flags().StringVar(&scope, "scope", "", "Scope of the field extraction rule. "+
		"This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from.")
	cmd.Flags().StringVar(&parseExpression, "parseExpression", "", "Specify the fields to be parsed")
	cmd.Flags().BoolVar(&enabled, "enabled", true, "Set to false if you don't want to enable the rule")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("scope")
	cmd.MarkFlagRequired("parseExpression")
	return cmd
}

func updateFieldExtractionRule(id string, name string, scope string, parseExpression string, enabled bool) {
	var fieldExtractionRulesResponse api.FieldExtractionRules
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateFieldExtractionRule{
		Name:            name,
		Scope:           scope,
		ParseExpression: parseExpression,
		Enabled:         enabled,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal request body")
	}
	requestUrl := "/v1/extractionRules/" + id
	client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}
	err = json.Unmarshal(responseBody, &fieldExtractionRulesResponse)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}
	fieldExtractionRulesResponseJson, err := json.MarshalIndent(fieldExtractionRulesResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("error marshalling response body")
	}
	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(fieldExtractionRulesResponseJson))
	}
}
