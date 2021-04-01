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
		merge           bool
		name            string
		sourceSyncMode  string
		timeZone        string
		targetCPU       int
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "updates a Sumo Logic collector settings",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.Flags().StringVar(&category, "category", "", "Specify a category for the collector")
	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Id of the collector you want to update")
	cmd.Flags().IntVar(&cutoffTimestamp, "cutoffTimestamp", 0, "Specify a cutoff timestamp for the collector, specified as milliseconds since epoch")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the collector")
	cmd.Flags().BoolVar(&ephemeral, "ephemeral", false, "When true the collector will be deleted after 12 hours of inactivity, defaults to false")
	cmd.Flags().StringVar(&fields, "fields", "", "Key value pair of fields (must be formatted as Key1:Value1,Key2:Value2)")
	cmd.Flags().StringVar(&hostName, "hostName", "", "Host name of the collector")
	cmd.Flags().BoolVar(&merge, "merge", true, "Merges the existing collector settings with the settings defined, set to true by default. If set to false it will overwrite the collector settings")
	cmd.Flags().StringVar(&name, "name", "", "Name of the collector, it must be unique on your account")
	cmd.Flags().StringVar(&sourceSyncMode, "sourceSyncMode", "", "For installed collectors whether the Collector is using local source of cloud management" +
		"(\"Json\" for local source and \"UI\" for cloud source this is only configurable on installed collectors")
	cmd.Flags().StringVar(&timeZone, "timeZone", "", "Time zone of the Collector. Refer to the TZ column of this site: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones")
	cmd.Flags().IntVar(&targetCPU, "targetCPU", 0, "When CPU utilisation exceeds this threshold, the Collector will slow down its rate of ingestion to lower its CPU utilisation" +
		"(only configurable on installable collectors)")
	cmd.MarkFlagRequired("collectorId")
	return cmd
}

func updateCollector(category string, collectorId int, cutoffTimestamp int, description string, ephemeral bool,
	fields string, hostName string, merge bool, name string, sourceSyncMode string, timeZone string, targetCPU int) {
	log := logging.GetConsoleLogger()
	var collectorInfo api.CollectorsResponse
	if merge == true {
		requestUrl := "v1/collectors/" + strconv.Itoa(collectorId)
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
		err = json.Unmarshal(responseBody, &collectorInfo)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		if response.StatusCode != 200 {
			log.Fatal().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
		}

		// Building body payload to update the collector based on the differences
		// between the current collector settings and the desired settings
		requestBodySchema := &api.CollectorsResponse{}
		if strings.EqualFold(collectorInfo.Category, category) {
			requestBodySchema.Category = collectorInfo.Category
		} else {
			requestBodySchema.Category = category
		}

		if collectorInfo.CutoffTimestamp == cutoffTimestamp {
			requestBodySchema.CutoffTimestamp = collectorInfo.CutoffTimestamp
		} else {
			requestBodySchema.CutoffTimestamp = cutoffTimestamp
		}

		if strings.EqualFold(collectorInfo.Description, description) {
			requestBodySchema.Description = collectorInfo.Description
		} else {
			requestBodySchema.Description = description
		}

		if collectorInfo.Ephemeral == ephemeral {
			requestBodySchema.Ephemeral = collectorInfo.Ephemeral
		} else {
			requestBodySchema.Ephemeral = ephemeral
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
		if reflect.DeepEqual(collectorInfo.Fields,fieldsMap) {
			requestBodySchema.Fields = collectorInfo.Fields
		} else {
			requestBodySchema.Fields = fieldsMap
		}

		if collectorInfo.HostName == hostName {
			requestBodySchema.HostName = collectorInfo.HostName
		} else {
			requestBodySchema.HostName = hostName
		}

		if collectorInfo.Name == name {
			requestBodySchema.Name = collectorInfo.Name
		} else {
			requestBodySchema.Name = name
		}

		if collectorInfo.CollectorType == "Installable" && collectorInfo.SourceSyncMode == sourceSyncMode {
			requestBodySchema.SourceSyncMode = collectorInfo.SourceSyncMode
		} else {
			requestBodySchema.SourceSyncMode = sourceSyncMode
		}

		if collectorInfo.TimeZone == timeZone {
			requestBodySchema.TimeZone = collectorInfo.TimeZone
		} else {
			requestBodySchema.TimeZone = timeZone
		}

		if collectorInfo.CollectorType == "Installable" && collectorInfo.TargetCpu == targetCPU {
			requestBodySchema.TargetCpu = collectorInfo.TargetCpu
		} else {
			requestBodySchema.TargetCpu = targetCPU
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
	} else {
		requestBodySchema := &api.
	}
}
