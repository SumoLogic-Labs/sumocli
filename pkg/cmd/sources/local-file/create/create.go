package create

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/api"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Incubator/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
	"strconv"
)

func NewCmdCreateLocalFileSource() *cobra.Command {
	var (
		automaticDateParsing       bool
		blacklist                  []string
		category                   string
		collectorId                int
		cutoffRelativeTime         string
		encoding                   string
		filterType                 string
		filterName                 string
		filterRegexp               string
		forceTimeZone              bool
		multilineProcessingEnabled bool
		name                       string
		pathExpression             string
		timeZone                   string
		useAutolineMatching        bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a local file source on the specified Sumo Logic collector",
		Long:  "Local file sources can only be created on installed collectors",
		Run: func(cmd *cobra.Command, args []string) {
			createLocalFileSource(automaticDateParsing, blacklist, category, collectorId, cutoffRelativeTime, encoding,
				filterType, filterName, filterRegexp, forceTimeZone, multilineProcessingEnabled, name, pathExpression,
				timeZone, useAutolineMatching)
		},
	}
	cmd.Flags().BoolVar(&automaticDateParsing, "automaticDateParsing", true, "Specify to false to prevent time stamp parsing")
	cmd.Flags().StringSliceVar(&blacklist, "blacklist", []string{}, "Specify a comma separated list of file exclusions")
	cmd.Flags().StringVar(&category, "category", "", "Specify the sourceCategory for the source")
	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Specify the collector id to associate the source to")
	cmd.Flags().StringVar(&cutoffRelativeTime, "cutoffRelativeTime", "-24h", "Specify a cutoff time stamp "+
		"for example: -1h (1 hour ago), -1d (1 day ago), -1w (1 week ago)")
	cmd.Flags().StringVar(&encoding, "encoding", "UTF-8", "Specify the encoding")
	cmd.Flags().StringVar(&filterType, "filterType", "", "Specify the filter type accepted values are "+
		"Exclude, Include, Hash, Mask or Forward")
	cmd.Flags().StringVar(&filterName, "filterName", "", "Specify the name of the rule")
	cmd.Flags().StringVar(&filterRegexp, "filterRegexp", "", "Specify a regular expression to define the filter")
	cmd.Flags().BoolVar(&forceTimeZone, "forceTimeZone", false, "Set to true to force the source to use a specific time zone")
	cmd.Flags().BoolVar(&multilineProcessingEnabled, "multilineProcessingEnabled", true, "Set to false to disable multiline processing")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the source")
	cmd.Flags().StringVar(&pathExpression, "pathExpression", "", "Specify the path to the log file(s)")
	cmd.Flags().StringVar(&timeZone, "timeZone", "UTC", "Specify the time zone the source should use")
	cmd.Flags().BoolVar(&useAutolineMatching, "useAutolineMatching", true, "Set to false to prevent message boundaries from being automatically inferred")
	cmd.MarkFlagRequired("category")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("pathExpression")
	return cmd
}

func createLocalFileSource(automaticDateParsing bool, blacklist []string, category string, collectorId int,
	cutoffRelativeTime string, encoding string, filterType string, filterName string, filterRegexp string,
	forceTimeZone bool, multilineProcessingEnabled bool, name string, pathExpression string, timeZone string,
	useAutolineMatching bool) {
	var sourceResponse api.GetLocalFileSource
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/collectors/" + strconv.Itoa(collectorId) + "/sources"
	requestBodySchema := &api.CreateLocalFileSource{
		ApiVersion: "v1",
		Source: api.LocalFileSource{
			Name:                       name,
			Category:                   category,
			AutomaticDateParsing:       automaticDateParsing,
			MultilineProcessingEnabled: multilineProcessingEnabled,
			UseAutolineMatching:        useAutolineMatching,
			ForceTimeZone:              forceTimeZone,
			TimeZone:                   timeZone,
			Filters:                    nil,
			CutoffRelativeTime:         cutoffRelativeTime,
			Encoding:                   encoding,
			PathExpression:             pathExpression,
			Blacklist:                  blacklist,
			SourceType:                 "LocalFile",
		},
	}
	if filterType != "" && filterName != "" && filterRegexp != "" {
		filter := api.SourceFilters{
			FilterType: filterType,
			Name:       filterName,
			Regexp:     filterRegexp,
		}
		requestBodySchema.Source.Filters = append(requestBodySchema.Source.Filters, filter)
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
