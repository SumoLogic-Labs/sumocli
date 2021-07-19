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

func NewCmdCollectorUpdate() *cobra.Command {
	var (
		category        string
		collectorId     int
		cutoffTimestamp int
		description     string
		ephemeral       bool
		fields          string
		hostName        string
		name            string
		sourceSyncMode  string
		timeZone        string
		targetCPU       int
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "updates a Sumo Logic collector settings",
		Run: func(cmd *cobra.Command, args []string) {
			updateCollector(category, collectorId, cutoffTimestamp, description, ephemeral, fields, hostName,
				name, sourceSyncMode, timeZone, targetCPU)
		},
	}
	cmd.Flags().StringVar(&category, "category", "", "Specify a category for the collector")
	cmd.Flags().IntVar(&collectorId, "id", 0, "Id of the collector you want to update")
	cmd.Flags().IntVar(&cutoffTimestamp, "cutoffTimestamp", 0, "Specify a cutoff timestamp for the collector, specified as milliseconds since epoch")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the collector")
	cmd.Flags().BoolVar(&ephemeral, "ephemeral", false, "When true the collector will be deleted after 12 hours of inactivity, defaults to false")
	cmd.Flags().StringVar(&fields, "fields", "", "Key value pair of fields (must be formatted as Key1:Value1,Key2:Value2)")
	cmd.Flags().StringVar(&hostName, "hostName", "", "Host name of the collector")
	cmd.Flags().StringVar(&name, "name", "", "Name of the collector, it must be unique on your account")
	cmd.Flags().StringVar(&sourceSyncMode, "sourceSyncMode", "", "For installed collectors whether the Collector is using local source of cloud management"+
		"(\"Json\" for local source and \"UI\" for cloud source this is only configurable on installed collectors")
	cmd.Flags().StringVar(&timeZone, "timeZone", "", "Time zone of the Collector. Refer to the TZ column of this site: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones")
	cmd.Flags().IntVar(&targetCPU, "targetCPU", 0, "When CPU utilisation exceeds this threshold, the Collector will slow down its rate of ingestion to lower its CPU utilisation"+
		"(only configurable on installable collectors)")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updateCollector(category string, collectorId int, cutoffTimestamp int, description string, ephemeral bool,
	fields string, hostName string, name string, sourceSyncMode string, timeZone string, targetCPU int) {
	log := logging.GetConsoleLogger()
	var collectorInfo api.CollectorResponse

	requestUrl := "/v1/collectors/" + strconv.Itoa(collectorId)
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
	err = json.Unmarshal(responseBody, &collectorInfo)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}
	if response.StatusCode != 200 {
		log.Fatal().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
	}

	// Building body payload to update the collector based on the differences
	// between the current collector settings and the desired settings
	requestBodySchema := &api.CollectorResponse{}
	if collectorInfo.Collector.Category == category {
		requestBodySchema.Collector.Category = collectorInfo.Collector.Category
	} else {
		requestBodySchema.Collector.Category = category
	}

	if collectorInfo.Collector.CutoffTimestamp == cutoffTimestamp {
		requestBodySchema.Collector.CutoffTimestamp = collectorInfo.Collector.CutoffTimestamp
	} else {
		requestBodySchema.Collector.CutoffTimestamp = cutoffTimestamp
	}

	if collectorInfo.Collector.Description == description {
		requestBodySchema.Collector.Description = collectorInfo.Collector.Description
	} else {
		requestBodySchema.Collector.Description = description
	}

	if collectorInfo.Collector.Ephemeral == ephemeral {
		requestBodySchema.Collector.Ephemeral = collectorInfo.Collector.Ephemeral
	} else {
		requestBodySchema.Collector.Ephemeral = ephemeral
	}

	fieldsMap := make(map[string]string)
	if fields != "" {
		splitStrings := strings.Split(fields, ",")
		for i, splitString := range splitStrings {
			components := strings.Split(splitString, ":")
			fieldsMap[components[0]] = components[1]
			i++
		}
	}
	if reflect.DeepEqual(collectorInfo.Collector.Fields, fieldsMap) {
		requestBodySchema.Collector.Fields = collectorInfo.Collector.Fields
	} else {
		requestBodySchema.Collector.Fields = fieldsMap
	}

	if collectorInfo.Collector.HostName == hostName {
		requestBodySchema.Collector.HostName = collectorInfo.Collector.HostName
	} else {
		requestBodySchema.Collector.HostName = hostName
	}

	if collectorInfo.Collector.Name == name {
		requestBodySchema.Collector.Name = collectorInfo.Collector.Name
	} else {
		requestBodySchema.Collector.Name = name
	}

	if collectorInfo.Collector.CollectorType == "Installable" && collectorInfo.Collector.SourceSyncMode == sourceSyncMode {
		requestBodySchema.Collector.SourceSyncMode = collectorInfo.Collector.SourceSyncMode
	} else {
		requestBodySchema.Collector.SourceSyncMode = sourceSyncMode
	}

	if collectorInfo.Collector.TimeZone == timeZone {
		requestBodySchema.Collector.TimeZone = collectorInfo.Collector.TimeZone
	} else {
		requestBodySchema.Collector.TimeZone = timeZone
	}

	if collectorInfo.Collector.CollectorType == "Installable" && collectorInfo.Collector.TargetCpu == targetCPU {
		requestBodySchema.Collector.TargetCpu = collectorInfo.Collector.TargetCpu
	} else {
		requestBodySchema.Collector.TargetCpu = targetCPU
	}
	requestBodySchema.Collector.Alive = collectorInfo.Collector.Alive
	requestBodySchema.Collector.CollectorType = collectorInfo.Collector.CollectorType
	requestBodySchema.Collector.CollectorVersion = collectorInfo.Collector.CollectorVersion
	requestBodySchema.Collector.CutoffRelativeTime = collectorInfo.Collector.CutoffRelativeTime
	requestBodySchema.Collector.Links = collectorInfo.Collector.Links
	requestBodySchema.Collector.Id = collectorInfo.Collector.Id
	requestBodySchema.Collector.LastSeenAlive = collectorInfo.Collector.LastSeenAlive
	if collectorInfo.Collector.CollectorType == "Installable" {
		requestBodySchema.Collector.OsArch = collectorInfo.Collector.OsArch
		requestBodySchema.Collector.OsName = collectorInfo.Collector.OsName
		requestBodySchema.Collector.OsTime = collectorInfo.Collector.OsTime
		requestBodySchema.Collector.OsVersion = collectorInfo.Collector.OsVersion
	}
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
	err = json.Unmarshal(responseBody, &collectorInfo)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}
	collectorInfoJson, err := json.MarshalIndent(&collectorInfo, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("error marshalling response body")
	}
	if response.StatusCode != 200 {
		log.Error().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
	} else {
		fmt.Println(string(collectorInfoJson))
	}
}
