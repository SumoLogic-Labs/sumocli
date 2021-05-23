package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"reflect"
	"strconv"
	"strings"
)

func NewCmdUpdateHttpSource() *cobra.Command {
	var (
		category                   string
		collectorId                int
		fieldNames                 string
		fieldValues                string
		messagePerRequest          bool
		multilineProcessingEnabled bool
		name                       string
		sourceId                   string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a specific HTTP source on the specified Sumo Logic collector",
		Run: func(cmd *cobra.Command, args []string) {
			updateHttpSource(category, collectorId, fieldNames, fieldValues, messagePerRequest,
				multilineProcessingEnabled, name, sourceId)
		},
	}
	cmd.Flags().StringVar(&category, "category", "", "Specify the sourceCategory for the source")
	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Specify the collector id the source is associated with")
	cmd.Flags().StringVar(&fieldNames, "fieldNames", "", "Specify the names of fields to add to the source "+
		"{names need to be comma separated e.g. field1,field2")
	cmd.Flags().StringVar(&fieldValues, "fieldValues", "", "Specify the values of fields to add to the source "+
		"(values need to be comma separated e.g. value1,value2")
	cmd.Flags().BoolVar(&messagePerRequest, "messagePerRequest", false, "Specify if there is one message per request")
	cmd.Flags().BoolVar(&multilineProcessingEnabled, "multilineProcessingEnabled", false, "Specify if multiline processing is enabled")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the source")
	cmd.Flags().StringVar(&sourceId, "sourceId", "", "Specify the source Id to update")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("name")
	return cmd
}

func updateHttpSource(category string, collectorId int, fieldNames string, fieldValues string,
	messagePerRequest bool, multilineProcessingEnabled bool, name string, sourceId string) {
	var sourceResponse api.CreateSourceResponse
	log := logging.GetConsoleLogger()

	requestUrl := "v1/collectors/" + strconv.Itoa(collectorId) + "/sources/" + sourceId
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
	etag := response.Header.Values("Etag")
	err = json.Unmarshal(responseBody, &sourceResponse)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}
	if response.StatusCode != 200 {
		log.Fatal().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
	}

	// Building body payload to update the source based on the differences
	// between the current source and the desired settings
	fieldsMap := make(map[string]string)
	if fieldNames != "" && fieldValues != "" {
		fieldNamesSlice := strings.Split(fieldNames, ",")
		fieldValuesSlice := strings.Split(fieldValues, ",")
		for i, _ := range fieldNamesSlice {
			fieldsMap[fieldNamesSlice[i]] = fieldValuesSlice[i]
			i++
		}
	}
	requestBodySchema := &api.CreateSourceResponse{}
	if sourceResponse.Source.Category == category {
		requestBodySchema.Source.Category = sourceResponse.Source.Category
	} else {
		requestBodySchema.Source.Category = category
	}

	if reflect.DeepEqual(sourceResponse.Source.Fields, fieldsMap) {
		requestBodySchema.Source.Fields = sourceResponse.Source.Fields
	} else {
		requestBodySchema.Source.Fields = fieldsMap
	}

	if sourceResponse.Source.MessagePerRequest == messagePerRequest {
		requestBodySchema.Source.MessagePerRequest = sourceResponse.Source.MessagePerRequest
	} else {
		requestBodySchema.Source.MessagePerRequest = messagePerRequest
	}

	if sourceResponse.Source.MultilineProcessingEnabled == multilineProcessingEnabled {
		requestBodySchema.Source.MultilineProcessingEnabled = sourceResponse.Source.MultilineProcessingEnabled
	} else {
		requestBodySchema.Source.MultilineProcessingEnabled = multilineProcessingEnabled
	}

	if sourceResponse.Source.Name == name {
		requestBodySchema.Source.Name = sourceResponse.Source.Name
	} else {
		requestBodySchema.Source.Name = name
	}
	requestBodySchema.Source.Id = sourceResponse.Source.Id
	requestBodySchema.Source.Alive = sourceResponse.Source.Alive
	requestBodySchema.Source.AutomaticDateParsing = sourceResponse.Source.AutomaticDateParsing
	requestBodySchema.Source.CutoffTimestamp = sourceResponse.Source.CutoffTimestamp
	requestBodySchema.Source.Encoding = sourceResponse.Source.Encoding
	requestBodySchema.Source.Filters = sourceResponse.Source.Filters
	requestBodySchema.Source.ForceTimezone = sourceResponse.Source.ForceTimezone
	requestBodySchema.Source.HostName = sourceResponse.Source.HostName
	requestBodySchema.Source.Interval = sourceResponse.Source.Interval
	requestBodySchema.Source.Metrics = sourceResponse.Source.Metrics
	requestBodySchema.Source.SourceType = sourceResponse.Source.SourceType
	requestBodySchema.Source.UseAutolineMatching = sourceResponse.Source.UseAutolineMatching
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal request body")
	}
	client, request = factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
	request.Header.Add("If-Match", etag[0])
	response, err = client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request " + requestUrl)
	}
	defer response.Body.Close()
	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}

	err = json.Unmarshal(responseBody, &sourceResponse)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}

	sourceResponseJson, err := json.MarshalIndent(&sourceResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("error marshalling response body")
	}

	if response.StatusCode != 200 {
		log.Error().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
	} else {
		fmt.Println(string(sourceResponseJson))
	}
}
