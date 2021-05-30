package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdDynamicParsingCreate() *cobra.Command {
	var (
		name    string
		scope   string
		enabled bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new dynamic parsing rule.",
		Run: func(cmd *cobra.Command, args []string) {
			createDynamicParsingRule(name, scope, enabled)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the dynamic parsing rule")
	cmd.Flags().StringVar(&scope, "scope", "", "Scope of the dynamic parsing rule. "+
		"This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from.")
	cmd.Flags().BoolVar(&enabled, "enabled", true, "Set to false if you don't want to enable the rule")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("scope")
	return cmd
}

func createDynamicParsingRule(name string, scope string, enabled bool) {
	var dynamicParsingRuleResponse api.DynamicParsingRules
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateDynamicParsingRule{
		Name:    name,
		Scope:   scope,
		Enabled: enabled,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/dynamicParsingRules"
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

	err = json.Unmarshal(responseBody, &dynamicParsingRuleResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	dynamicParsingRuleResponseJson, err := json.MarshalIndent(dynamicParsingRuleResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(dynamicParsingRuleResponseJson))
	}
}
