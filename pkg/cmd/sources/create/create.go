package create

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io/ioutil"
	"strconv"
	"strings"
)

func NewCmdCreateSource() *cobra.Command {
	var (
		category                   string
		collectorId                int
		fields                     string
		messagePerRequest          bool
		multilineProcessingEnabled bool
		name                       string
		httpSource                 bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a source on the specified Sumo Logic collector",
		Run: func(cmd *cobra.Command, args []string) {
			var fieldsMap map[string]string
			log := logging.GetConsoleLogger()
			if fields != "" {
				fieldsMap = make(map[string]string)
				splitStrings := strings.Split(fields, ",")
				for i, splitString := range splitStrings {
					components := strings.Split(splitString, ":")
					fieldsMap[components[0]] = components[1]
					i++
				}
			}
			if httpSource == true {
				HTTPSource(category, fieldsMap, messagePerRequest, multilineProcessingEnabled, name, collectorId, log)
			}
		},
	}

	cmd.Flags().StringVar(&category, "category", "", "Specify the sourceCategory for the source")
	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Specify the collector id to attach the source to")
	cmd.Flags().StringVar(&fields, "fields", "", "Specify fields to add to the source")
	cmd.Flags().BoolVar(&httpSource, "httpSource", false, "Specifies a HTTP source for creation")
	cmd.Flags().BoolVar(&messagePerRequest, "messagePerRequest", false, "Specify if there is one message per request")
	cmd.Flags().BoolVar(&multilineProcessingEnabled, "multilineProcessingEnabled", false, "Specify if multiline processing is enabled")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the source")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("name")
	return cmd
}

func HTTPSource(category string, fields map[string]string, messagePerRequest bool, multilineProcessingEnabled bool,
	name string, collectorId int, log zerolog.Logger) api.CreateSourceResponse {
	var sourceResponse api.CreateSourceResponse
	requestUrl := "v1/collectors/" + strconv.Itoa(collectorId) + "/sources"
	requestBodySchema := &api.CreateHTTPSource{
		ApiVersion: "",
		Source: api.HttpSource{
			SourceType:                 "HTTP",
			Name:                       name,
			Category:                   category,
			Fields:                     fields,
			MessagePerRequest:          messagePerRequest,
			MultilineProcessingEnabled: multilineProcessingEnabled,
		},
	}
	requestBody, _ := json.Marshal(requestBodySchema)
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}

	jsonErr := json.Unmarshal(responseBody, &sourceResponse)
	if jsonErr != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}
	return sourceResponse
}
