package create

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

func NewCmdCreateHttpSource() *cobra.Command {
	var (
		category                   string
		collectorId                int
		fieldNames                 string
		fieldValues                string
		messagePerRequest          bool
		multilineProcessingEnabled bool
		name                       string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a HTTP source on the specified Sumo Logic collector",
		Run: func(cmd *cobra.Command, args []string) {
			createHttpSource(category, collectorId, fieldNames, fieldValues, messagePerRequest,
				multilineProcessingEnabled, name)
		},
	}
	cmd.Flags().StringVar(&category, "category", "", "Specify the sourceCategory for the source")
	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Specify the collector id to associate the source to")
	cmd.Flags().StringVar(&fieldNames, "fieldNames", "", "Specify the names of fields to add to the source "+
		"{names need to be comma separated e.g. field1,field2")
	cmd.Flags().StringVar(&fieldValues, "fieldValues", "", "Specify the values of fields to add to the source "+
		"(values need to be comma separated e.g. value1,value2")
	cmd.Flags().BoolVar(&messagePerRequest, "messagePerRequest", false, "Specify if there is one message per request")
	cmd.Flags().BoolVar(&multilineProcessingEnabled, "multilineProcessingEnabled", false, "Specify if multiline processing is enabled")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the source")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createHttpSource(category string, collectorId int, fieldNames string, fieldValues string,
	messagePerRequest bool, multilineProcessingEnabled bool, name string) {
	var sourceResponse api.CreateSourceResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v1/collectors/" + strconv.Itoa(collectorId) + "/sources"
	fieldsMap := make(map[string]string)
	if fieldNames != "" && fieldValues != "" {
		fieldNamesSlice := strings.Split(fieldNames, ",")
		fieldValuesSlice := strings.Split(fieldValues, ",")
		for i, _ := range fieldNamesSlice {
			fieldsMap[fieldNamesSlice[i]] = fieldValuesSlice[i]
			i++
		}
	}
	requestBodySchema := &api.CreateHTTPSource{
		ApiVersion: "v1",
		Source: api.HttpSource{
			SourceType:                 "HTTP",
			Name:                       name,
			Category:                   category,
			Fields:                     fieldsMap,
			MessagePerRequest:          messagePerRequest,
			MultilineProcessingEnabled: multilineProcessingEnabled,
		},
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal request body")
	}
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}

	err = json.Unmarshal(responseBody, &sourceResponse)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}

	sourceResponseJson, err := json.MarshalIndent(sourceResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 201 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(sourceResponseJson))
	}
}
