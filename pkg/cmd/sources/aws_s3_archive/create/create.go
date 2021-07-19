package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/internal/cloud"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"strconv"
	"strings"
)

func NewCmdAWSS3ArchiveSourceCreate() *cobra.Command {
	var (
		automaticDateParsing       bool
		awsId                      string
		awsKey                     string
		category                   string
		collectorId                int
		cutoffRelativeTime         string
		encoding                   string
		fieldNames                 string
		fieldValues                string
		filterType                 string
		filterName                 string
		filterRegexp               string
		forceTimeZone              bool
		iamRoleArn                 string
		multilineProcessingEnabled bool
		name                       string
		pathExpression             string
		paused                     bool
		s3BucketName               string
		scanInterval               int
		timeZone                   string
		useAutolineMatching        bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Configures AWS S3 Archives",
		Long: "AWS S3 Archive sources allow you to ingest data from an S3 archive. " +
			"Once the source is created you can start an ingestion job by running " +
			"sumocli archive-ingestion create." +
			"Note: You will need to ensure that the access method you use for Sumo Logic has the ability to list and " +
			"get objects from the specified S3 bucket.",
		Run: func(cmd *cobra.Command, args []string) {
			createAWSS3ArchiveSource(automaticDateParsing, awsId, awsKey, category, collectorId, cutoffRelativeTime, encoding,
				fieldNames, fieldValues, filterType, filterName, filterRegexp, forceTimeZone, iamRoleArn, multilineProcessingEnabled,
				name, pathExpression, paused, s3BucketName, scanInterval, timeZone, useAutolineMatching)
		},
	}
	cmd.Flags().BoolVar(&automaticDateParsing, "automaticDateParsing", true, "Set to false if you don't want automatic date parsing")
	cmd.Flags().StringVar(&awsId, "awsId", "", "Specify an AWS ID that will be used to allow Sumo Logic "+
		"to collect data from AWS. It is preferred to use an IAM Role for access.")
	cmd.Flags().StringVar(&awsKey, "awsKey", "", "Specify an AWS Key that will be used to allow Sumo Logic "+
		"to collect data from AWS. It is preferred to use an IAM Role for access.")
	cmd.Flags().StringVar(&category, "category", "", "Specify the sourceCategory for the source")
	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Specify the hosted collectorId to attach the "+
		"source to.")
	cmd.Flags().StringVar(&cutoffRelativeTime, "cutoffRelativeTime", "-24h", "Specify a cutoff time stamp "+
		"for example: -1h (1 hour ago), -1d (1 day ago), -1w (1 week ago)")
	cmd.Flags().StringVar(&encoding, "encoding", "UTF-8", "Specify the encoding")
	cmd.Flags().StringVar(&fieldNames, "fieldNames", "", "Specify the names of fields to add to the source "+
		"{names need to be comma separated e.g. field1,field2")
	cmd.Flags().StringVar(&fieldValues, "fieldValues", "", "Specify the values of fields to add to the source "+
		"(values need to be comma separated e.g. value1,value2")
	cmd.Flags().StringVar(&filterType, "filterType", "", "Specify the filter type accepted values are "+
		"Exclude, Include, Hash, Mask or Forward")
	cmd.Flags().StringVar(&filterName, "filterName", "", "Specify the name of the rule")
	cmd.Flags().StringVar(&filterRegexp, "filterRegexp", "", "Specify a regular expression to define the filter")
	cmd.Flags().BoolVar(&forceTimeZone, "forceTimeZone", false, "Set to true to force the source to use a specific time zone")
	cmd.Flags().StringVar(&iamRoleArn, "iamRoleArn", "", "Specify an IAM role arn for Sumo Logic to use "+
		"to collect data from AWS. This is the preferred method of authentication.")
	cmd.Flags().BoolVar(&multilineProcessingEnabled, "multilineProcessingEnabled", true, "Set to false to disable multiline processing")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the source")
	cmd.Flags().StringVar(&pathExpression, "pathExpression", "", "Specify the path to the log files "+
		"in the s3 bucket")
	cmd.Flags().BoolVar(&paused, "paused", false, "Set to true if you want to pause collection")
	cmd.Flags().StringVar(&s3BucketName, "s3BucketName", "", "Specify the s3 bucket name (if createAWSInfrastructure is "+
		"set to true the bucket name needs to be globally unique)")
	cmd.Flags().IntVar(&scanInterval, "scanInterval", 300000, "Specify the time interval of S3 bucket scans "+
		"for new data, in milliseconds. Minimum value is 1000. For automatic set value to -1.")
	cmd.Flags().StringVar(&timeZone, "timeZone", "UTC", "Specify the time zone the source should use")
	cmd.Flags().BoolVar(&useAutolineMatching, "useAutolineMatching", true, "Set to false to prevent message boundaries from being automatically inferred")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("pathExpression")
	cmd.MarkFlagRequired("s3BucketName")
	return cmd
}

func createAWSS3ArchiveSource(automaticDateParsing bool, awsId string, awsKey string, category string,
	collectorId int, cutoffRelativeTime string, encoding string, fieldNames string, fieldValues string, filterType string,
	filterName string, filterRegexp string, forceTimeZone bool, iamRoleArn string, multilineProcessingEnabled bool, name string,
	pathExpression string, paused bool, s3BucketName string, scanInterval int, timeZone string, useAutolineMatching bool) {
	var s3ArchiveResponse api.AWSCloudTrailCollectionResponse
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/collectors/" + strconv.Itoa(collectorId) + "/sources"
	fieldsMap := make(map[string]string)
	if fieldNames != "" && fieldValues != "" {
		fieldNamesSlice := strings.Split(fieldNames, ",")
		fieldValuesSlice := strings.Split(fieldValues, ",")
		for i, _ := range fieldNamesSlice {
			fieldsMap[fieldNamesSlice[i]] = fieldValuesSlice[i]
			i++
		}
	}
	requestBodySchema := &api.AWSCloudTrailCollection{
		ApiVersion: "",
		Source: api.AWSCloudTrail{
			SourceType:                 "Polling",
			Name:                       name,
			Category:                   category,
			ContentType:                "AwsS3ArchiveBucket",
			ScanInterval:               scanInterval,
			Paused:                     paused,
			AutomaticDateParsing:       automaticDateParsing,
			MultilineProcessingEnabled: multilineProcessingEnabled,
			UseAutolineMatching:        useAutolineMatching,
			ForceTimeZone:              forceTimeZone,
			TimeZone:                   timeZone,
			Filters:                    nil,
			CutoffRelativeTime:         cutoffRelativeTime,
			Encoding:                   encoding,
			Fields:                     fieldsMap,
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
	thirdPartyReference := api.ThirdPartyReferenceResources{
		ServiceType: "AwsS3ArchiveBucket",
		Path: api.ThirdPartyReferenceResourcesPath{
			Type:           "S3BucketPathExpression",
			BucketName:     s3BucketName,
			PathExpression: pathExpression,
		},
	}
	thirdPartyReference.Authentication = cloud.SetAWSAuthentication(awsId, awsKey, iamRoleArn)
	requestBodySchema.Source.ThirdPartyRef.Resources = append(requestBodySchema.Source.ThirdPartyRef.Resources, thirdPartyReference)

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

	err = json.Unmarshal(responseBody, &s3ArchiveResponse)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}

	s3ArchiveResponseJson, err := json.MarshalIndent(s3ArchiveResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 201 {
		log.Error().Msg("failed to create source " + string(responseBody))
	} else {
		fmt.Println(string(s3ArchiveResponseJson))
	}
}
