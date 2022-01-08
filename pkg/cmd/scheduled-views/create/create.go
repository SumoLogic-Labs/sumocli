package create

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/api"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Labs/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdScheduledViewsCreate() *cobra.Command {
	var (
		query            string
		indexName        string
		startTime        string
		retentionPeriod  int
		dataForwardingId string
		parsingMode      string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new scheduled view in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			createScheduledView(query, indexName, startTime, retentionPeriod, dataForwardingId, parsingMode)
		},
	}
	cmd.Flags().StringVar(&query, "query", "", "Specify the query that defines the data to be included in the scheduled view")
	cmd.Flags().StringVar(&indexName, "indexName", "", "Specify the name of the index for the scheduled view")
	cmd.Flags().StringVar(&startTime, "startTime", "", "Specify a start timestamp in UTC in RFC3339 format. "+
		"Example: 2021-05-30T12:00:00Z")
	cmd.Flags().IntVar(&retentionPeriod, "retentionPeriod", -1, "Specify the number of days to retain data in the partition. "+
		"-1 specifies that the default value for the account is used.")
	cmd.Flags().StringVar(&dataForwardingId, "dataForwardingId", "", "Specify an ID of a data forwarding configuration to be used by the scheduled view")
	cmd.Flags().StringVar(&parsingMode, "parsingMode", "Manual", "Specify the parsing mode to scan the JSON format log messages. "+
		"Possible values are AutoParse and Manual")
	cmd.MarkFlagRequired("query")
	cmd.MarkFlagRequired("indexName")
	cmd.MarkFlagRequired("startTime")
	return cmd
}

func createScheduledView(query string, indexName string, startTime string, retentionPeriod int, dataForwardingId string,
	parsingMode string) {
	var scheduledViewsResponse api.ScheduledViews
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateScheduledView{
		Query:            query,
		IndexName:        indexName,
		StartTime:        startTime,
		RetentionPeriod:  retentionPeriod,
		DataForwardingId: dataForwardingId,
		ParsingMode:      parsingMode,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "/v1/scheduledViews"
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

	err = json.Unmarshal(responseBody, &scheduledViewsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	scheduledViewsResponseJson, err := json.MarshalIndent(scheduledViewsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(scheduledViewsResponseJson))
	}
}
