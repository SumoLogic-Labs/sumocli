package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"strconv"
	"strings"
)

func NewCmdDynamicParsingUpdate() *cobra.Command {
	var (
		id      string
		name    string
		scope   string
		enabled bool
		merge   bool
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing dynamic parsing rule",
		Run: func(cmd *cobra.Command, args []string) {
			updateDynamicParsingRule(id, name, scope, enabled, merge)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the dynamic parsing rule")
	cmd.Flags().BoolVar(&merge, "merge", true, "If set to false it will overwrite the dynamic parsing rule configuration")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the dynamic parsing rule")
	cmd.Flags().StringVar(&scope, "scope", "", "Scope of the dynamic parsing rule. "+
		"This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from.")
	cmd.Flags().BoolVar(&enabled, "enabled", true, "Set to false if you don't want to enable the rule")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("scope")
	return cmd
}

func updateDynamicParsingRule(id string, name string, scope string, enabled bool, merge bool) {
	var dynamicParsingRuleResponse api.DynamicParsingRules
	log := logging.GetConsoleLogger()
	if merge == true {
		requestUrl := "/v1/dynamicParsingRules/" + id
		client, request := factory.NewHttpRequest("GET", requestUrl)
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request " + requestUrl)
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}
		err = json.Unmarshal(responseBody, &dynamicParsingRuleResponse)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		if response.StatusCode != 200 {
			log.Fatal().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
		}

		// Building body payload to update the ingest budget based on the differences
		// between the current ingest budget and the desired settings
		requestBodySchema := &api.CreateDynamicParsingRule{}
		if strings.EqualFold(dynamicParsingRuleResponse.Name, name) {
			requestBodySchema.Name = dynamicParsingRuleResponse.Name
		} else {
			requestBodySchema.Name = name
		}

		if strings.EqualFold(dynamicParsingRuleResponse.Scope, scope) {
			requestBodySchema.Scope = dynamicParsingRuleResponse.Scope
		} else {
			requestBodySchema.Scope = scope
		}

		if dynamicParsingRuleResponse.Enabled == enabled {
			requestBodySchema.Enabled = dynamicParsingRuleResponse.Enabled
		} else {
			requestBodySchema.Enabled = enabled
		}

		requestBody, err := json.Marshal(requestBodySchema)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request = factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
		response, err = client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request " + requestUrl)
		}
		defer response.Body.Close()
		responseBody, err = io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}
		err = json.Unmarshal(responseBody, &dynamicParsingRuleResponse)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		dynamicParsingRuleResponseJson, err := json.MarshalIndent(dynamicParsingRuleResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("error marshalling response body")
		}
		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, log)
		} else {
			fmt.Println(string(dynamicParsingRuleResponseJson))
		}
	} else {
		requestBodySchema := &api.CreateDynamicParsingRule{
			Name:    name,
			Scope:   scope,
			Enabled: enabled,
		}
		requestBody, err := json.Marshal(requestBodySchema)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		requestUrl := "/v1/dynamicParsingRules/" + id
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
		err = json.Unmarshal(responseBody, &dynamicParsingRuleResponse)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		dynamicParsingRuleResponseJson, err := json.MarshalIndent(dynamicParsingRuleResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("error marshalling response body")
		}
		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, log)
		} else {
			fmt.Println(string(dynamicParsingRuleResponseJson))
		}
	}
}
