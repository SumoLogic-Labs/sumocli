package create

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/SumoLogic-Labs/sumocli/api"
	"github.com/SumoLogic-Labs/sumocli/internal/cloud"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Labs/sumocli/pkg/logging"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/cobra"
	"io"
	"strconv"
	"strings"
)

func NewCmdAWSCloudTrailSourceCreate() *cobra.Command {
	var (
		automaticDateParsing       bool
		awsId                      string
		awsKey                     string
		awsRegion                  string
		category                   string
		cloudTrailName             string
		createAWSInfrastructure    bool
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
		snsTopicName               string
		snsTopicArn                string
		scanInterval               int
		timeZone                   string
		useAutolineMatching        bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Configures AWS CloudTrail log collection",
		Long: "AWS CloudTrail records API calls made to AWS. " +
			"This includes calls made using the AWS Management Console, AWS SDKs, command line tools, and higher-level AWS services. " +
			"You can use the --createAWSInfrastructure argument to deploy all the required AWS Infrastructure." +
			"This command will configure SNS Topic notifications with Sumo Logic if the --createAWSInfrastructure argument is used " +
			"or if --snsTopicArn is argument is used." +
			"Note: The AWS CloudTrail Source automatically parses the logs prior to upload.",
		Run: func(cmd *cobra.Command, args []string) {
			createAWSCloudTrailSource(automaticDateParsing, awsId, awsKey, awsRegion, category, cloudTrailName, createAWSInfrastructure,
				collectorId, cutoffRelativeTime, encoding, fieldNames, fieldValues, filterType, filterName, filterRegexp, forceTimeZone, iamRoleArn,
				multilineProcessingEnabled, name, pathExpression, paused, s3BucketName, snsTopicName, snsTopicArn,
				scanInterval, timeZone, useAutolineMatching)
		},
	}
	cmd.Flags().BoolVar(&automaticDateParsing, "automaticDateParsing", true, "Set to false if you don't want automatic date parsing")
	cmd.Flags().StringVar(&awsId, "awsId", "", "Specify an AWS ID that will be used to allow Sumo Logic "+
		"to collect data from AWS. It is preferred to use an IAM Role for access.")
	cmd.Flags().StringVar(&awsKey, "awsKey", "", "Specify an AWS Key that will be used to allow Sumo Logic "+
		"to collect data from AWS. It is preferred to use an IAM Role for access.")
	cmd.Flags().StringVar(&awsRegion, "awsRegion", "", "Specify an AWS region to deploy resources to "+
		"(this is only needed if createAWSInfrastructure is set to true)")
	cmd.Flags().StringVar(&category, "category", "", "Specify the sourceCategory for the source")
	cmd.Flags().StringVar(&cloudTrailName, "cloudTrailName", "", "Specify the CloudTrail name "+
		"(this is only needed if createAWSInfrastructure is set to true)")
	cmd.Flags().BoolVar(&createAWSInfrastructure, "createAWSInfrastructure", false, "Set to true if "+
		"you would like sumocli to create AWS CloudTrail infrastructure.")
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
	cmd.Flags().StringVar(&snsTopicName, "snsTopicName", "", "Specify a name for the SNS topic "+
		"(this is only needed if createAWSInfrastructure is set to true)")
	cmd.Flags().StringVar(&snsTopicArn, "snsTopicArn", "", "Specify a SNS topic ARN")
	cmd.Flags().IntVar(&scanInterval, "scanInterval", 300000, "Specify the time interval of S3 bucket scans "+
		"for new data, in milliseconds. Minimum value is 1000. For automatic set value to -1.")
	cmd.Flags().StringVar(&timeZone, "timeZone", "UTC", "Specify the time zone the source should use")
	cmd.Flags().BoolVar(&useAutolineMatching, "useAutolineMatching", true, "Set to false to prevent message boundaries from being automatically inferred")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("pathExpression")
	cmd.MarkFlagRequired("s3BucketName")
	return cmd
}

func createAWSCloudTrailSource(automaticDateParsing bool, awsId string, awsKey string, awsRegion string, category string, cloudTrailName string,
	createAWSInfrastructure bool, collectorId int, cutoffRelativeTime string, encoding string, fieldNames string, fieldValues string,
	filterType string, filterName string, filterRegexp string, forceTimeZone bool, iamRoleArn string, multilineProcessingEnabled bool,
	name string, pathExpression string, paused bool, s3BucketName string, snsTopicName string, snsTopicArn string,
	scanInterval int, timeZone string, useAutolineMatching bool) {
	var awsCloudTrailResponse api.AWSCloudTrailCollectionResponse
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/collectors/" + strconv.Itoa(collectorId) + "/sources"
	s3BucketExistence := false
	kmsKeyExistence := false
	kmsKeyId := ""
	var kmsKey *kms.CreateKeyOutput
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion(awsRegion))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load AWS config")
	}
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
		ApiVersion: "v1",
		Source: api.AWSCloudTrail{
			SourceType:                 "Polling",
			Name:                       name,
			Category:                   category,
			ContentType:                "AwsCloudTrailBucket",
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
	if createAWSInfrastructure == true {
		if cloudTrailName == "" {
			log.Error().Msg("please specify a cloudTrailName")
		} else if snsTopicName == "" {
			log.Error().Msg("please specify a snsTopicName")
		}
		stsClient := sts.NewFromConfig(cfg)
		callerInfo, err := stsClient.GetCallerIdentity(context.TODO(),
			&sts.GetCallerIdentityInput{})

		snsPolicy := api.AWSSNSPolicy{
			Version:   "2012-10-17",
			Statement: nil,
		}
		snsPolicyStatement := api.AWSSNSPolicyStatement{
			Sid:    "Allow S3 Access",
			Effect: "Allow",
			Principal: api.AWSSNSPolicyPrincipal{
				Service: "s3.amazonaws.com",
			},
			Action:   []string{"sns:Publish"},
			Resource: []string{"arn:aws:sns:" + awsRegion + ":" + *callerInfo.Account + ":" + snsTopicName},
			Condition: api.AWSSNSPolicyCondition{
				ArnLike: api.AWSSNSSourceArn{
					AWSSourceArn: "arn:aws:s3:*:*:" + s3BucketName,
				},
				StringEquals: api.AWSSNSSourceAccount{
					AWSSourceAccount: to.String(callerInfo.Account),
				},
			},
		}
		snsPolicy.Statement = append(snsPolicy.Statement, snsPolicyStatement)
		snsPolicyRequest, err := json.Marshal(snsPolicy)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal sns policy")
		}
		snsClient := sns.NewFromConfig(cfg)
		snsTopic, err := snsClient.CreateTopic(context.TODO(),
			&sns.CreateTopicInput{
				Name: to.StringPtr(snsTopicName),
				Attributes: map[string]string{
					"Policy": string(snsPolicyRequest),
				},
			})
		if err != nil {
			log.Error().Err(err).Msg("failed to create SNS topic")
		}
		snsTopicArn = *snsTopic.TopicArn

		s3Client := s3.NewFromConfig(cfg)
		buckets, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
		for _, bucket := range buckets.Buckets {
			if *bucket.Name == s3BucketName {
				s3BucketExistence = true
			}
		}
		if s3BucketExistence == false {
			_, err = s3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
				Bucket: to.StringPtr(s3BucketName),
				CreateBucketConfiguration: &types.CreateBucketConfiguration{
					LocationConstraint: "ap-southeast-2",
				},
			})
			if err != nil {
				log.Error().Err(err).Msg("failed to create s3 bucket")
			}
			s3BucketPolicy := api.AWSPolicy{
				Version:   "2012-10-17",
				Statement: nil,
			}
			cloudTrailAclCheck := api.AWSPolicyStatement{
				Sid:    "AWSCloudTrailAclCheck",
				Effect: "Allow",
				Principal: api.AWSPolicyPrincipal{
					Service: "cloudtrail.amazonaws.com",
				},
				Action:   []string{"s3:GetBucketAcl"},
				Resource: []string{"arn:aws:s3:::" + s3BucketName},
			}
			s3BucketPolicy.Statement = append(s3BucketPolicy.Statement, cloudTrailAclCheck)
			cloudTrailWrite := api.AWSPolicyStatement{
				Sid:    "AWSCloudTrailWrite",
				Effect: "Allow",
				Principal: api.AWSPolicyPrincipal{
					Service: "cloudtrail.amazonaws.com",
				},
				Action:   []string{"s3:PutObject"},
				Resource: []string{"arn:aws:s3:::" + s3BucketName + "/*"},
				Condition: api.AWSPolicyCondition{
					StringEquals: api.S3XAmzAcl{
						S3XAmzAcl: "bucket-owner-full-control",
					},
				},
			}
			s3BucketPolicy.Statement = append(s3BucketPolicy.Statement, cloudTrailWrite)
			s3BucketPolicyRequest, err := json.Marshal(s3BucketPolicy)
			if err != nil {
				log.Error().Err(err).Msg("failed to marshal cloudtrail bucket policy")
			}
			_, err = s3Client.PutBucketPolicy(context.TODO(), &s3.PutBucketPolicyInput{
				Bucket: to.StringPtr(s3BucketName),
				Policy: to.StringPtr(string(s3BucketPolicyRequest)),
			})
			if err != nil {
				log.Error().Err(err).Msg("failed to attach S3 bucket policy")
			}
			var snsTopicConfiguration []types.TopicConfiguration
			snsTopicConfigurationItem := types.TopicConfiguration{
				Events: []types.Event{
					"s3:ObjectCreated:*",
				},
				TopicArn: snsTopic.TopicArn,
				Filter:   nil,
			}
			snsTopicConfiguration = append(snsTopicConfiguration, snsTopicConfigurationItem)
			_, err = s3Client.PutBucketNotificationConfiguration(context.TODO(), &s3.PutBucketNotificationConfigurationInput{
				Bucket: to.StringPtr(s3BucketName),
				NotificationConfiguration: &types.NotificationConfiguration{
					TopicConfigurations: snsTopicConfiguration,
				},
				ExpectedBucketOwner: callerInfo.Account,
			})
			if err != nil {
				log.Error().Err(err).Msg("failed to created bucket notification configuration")
			}
		}

		kmsClient := kms.NewFromConfig(cfg)
		kmsKeyAliases, err := kmsClient.ListAliases(context.TODO(), &kms.ListAliasesInput{})
		if len(kmsKeyAliases.Aliases) > 0 {
			for _, kmsAlias := range kmsKeyAliases.Aliases {
				if *kmsAlias.AliasName == "alias/sumocli/"+cloudTrailName {
					kmsKeyExistence = true
					kmsKeyId = *kmsAlias.TargetKeyId
				}
			}
		}
		if kmsKeyExistence == false {
			kmsKey, err = kmsClient.CreateKey(context.TODO(),
				&kms.CreateKeyInput{
					Description: to.StringPtr("KMS Key used to encrypt CloudTrail provisioned by sumocli."),
					KeyUsage:    "ENCRYPT_DECRYPT",
				})
			if err != nil {
				log.Error().Err(err).Msg("failed to create kms key")
			}
			_, err = kmsClient.CreateAlias(context.TODO(), &kms.CreateAliasInput{
				AliasName:   to.StringPtr("alias/sumocli/" + cloudTrailName),
				TargetKeyId: kmsKey.KeyMetadata.KeyId,
			})
			if err != nil {
				log.Error().Err(err).Msg("failed to create kms key alias")
			}

			kmsPolicy := api.AWSPolicy{
				Version:   "2012-10-17",
				Statement: nil,
			}
			kmsCloudTrailCallerAccess := api.AWSPolicyStatement{
				Sid:    "Allow administrator access",
				Effect: "Allow",
				Principal: api.AWSPolicyPrincipal{
					AWS: []string{*callerInfo.Arn},
				},
				Action: []string{
					"kms:Create*",
					"kms:Describe*",
					"kms:Enable*",
					"kms:List*",
					"kms:Put*",
					"kms:Update*",
					"kms:Revoke*",
					"kms:Disable*",
					"kms:Get*",
					"kms:Delete*",
					"kms:ScheduleKeyDeletion",
					"kms:CancelKeyDeletion",
				},
				Resource: []string{"*"},
			}
			kmsPolicy.Statement = append(kmsPolicy.Statement, kmsCloudTrailCallerAccess)
			kmsCloudTrailEncrypt := api.AWSPolicyStatement{
				Sid:    "Allow CloudTrail to encrypt logs",
				Effect: "Allow",
				Principal: api.AWSPolicyPrincipal{
					Service: "cloudtrail.amazonaws.com",
				},
				Action:   []string{"kms:GenerateDataKey*"},
				Resource: []string{"*"},
				Condition: api.AWSPolicyCondition{
					StringLike: api.KMSEncryptionContext{
						KMSEncryption: []string{"arn:aws:cloudtrail:*:" + *callerInfo.Account + ":trail/*"},
					},
				},
			}
			kmsPolicy.Statement = append(kmsPolicy.Statement, kmsCloudTrailEncrypt)
			kmsCloudTrailDecrypt := api.AWSPolicyStatement{
				Sid:    "Enable encrypted CloudTrail log read access",
				Effect: "Allow",
				Principal: api.AWSPolicyPrincipal{
					AWS: []string{iamRoleArn},
				},
				Action:   []string{"kms:Decrypt"},
				Resource: []string{"*"},
				Condition: api.AWSPolicyCondition{
					Null: api.KMSEncryptionContextBool{
						KMSEncryption: false,
					},
				},
			}
			kmsPolicy.Statement = append(kmsPolicy.Statement, kmsCloudTrailDecrypt)
			kmsCloudTrailDescribe := api.AWSPolicyStatement{
				Sid:    "Allow CloudTrail access",
				Effect: "Allow",
				Principal: api.AWSPolicyPrincipal{
					Service: "cloudtrail.amazonaws.com",
				},
				Action:   []string{"kms:DescribeKey"},
				Resource: []string{"*"},
			}
			kmsPolicy.Statement = append(kmsPolicy.Statement, kmsCloudTrailDescribe)
			kmsPolicyRequest, err := json.Marshal(kmsPolicy)
			if err != nil {
				log.Error().Err(err).Msg("failed to marshal kms policy")
			}
			_, err = kmsClient.PutKeyPolicy(context.TODO(),
				&kms.PutKeyPolicyInput{
					KeyId:      kmsKey.KeyMetadata.KeyId,
					Policy:     to.StringPtr(string(kmsPolicyRequest)),
					PolicyName: to.StringPtr("default"),
				})
			if err != nil {
				log.Error().Err(err).Msg("failed to attach kms key policy")
			}
		}
		cloudtrailClient := cloudtrail.NewFromConfig(cfg)
		_, err = cloudtrailClient.GetTrail(context.TODO(), &cloudtrail.GetTrailInput{
			Name: to.StringPtr(cloudTrailName),
		})
		if err != nil {
			if kmsKeyExistence == true {
				trail, err := cloudtrailClient.CreateTrail(
					context.TODO(), &cloudtrail.CreateTrailInput{
						Name:                       to.StringPtr(cloudTrailName),
						S3BucketName:               to.StringPtr(s3BucketName),
						EnableLogFileValidation:    to.BoolPtr(true),
						IncludeGlobalServiceEvents: to.BoolPtr(true),
						IsMultiRegionTrail:         to.BoolPtr(true),
						IsOrganizationTrail:        to.BoolPtr(false),
						KmsKeyId:                   to.StringPtr(kmsKeyId),
					})
				if err != nil {
					log.Error().Err(err).Msg("failed to create AWS CloudTrail")
				}
				_, err = cloudtrailClient.StartLogging(context.TODO(),
					&cloudtrail.StartLoggingInput{
						Name: trail.TrailARN,
					})
				if err != nil {
					log.Error().Err(err).Msg("failed to start trail")
				}
			} else {
				trail, err := cloudtrailClient.CreateTrail(
					context.TODO(), &cloudtrail.CreateTrailInput{
						Name:                       to.StringPtr(cloudTrailName),
						S3BucketName:               to.StringPtr(s3BucketName),
						EnableLogFileValidation:    to.BoolPtr(true),
						IncludeGlobalServiceEvents: to.BoolPtr(true),
						IsMultiRegionTrail:         to.BoolPtr(true),
						IsOrganizationTrail:        to.BoolPtr(false),
						KmsKeyId:                   kmsKey.KeyMetadata.KeyId,
					})
				if err != nil {
					log.Error().Err(err).Msg("failed to create AWS CloudTrail")
				}
				_, err = cloudtrailClient.StartLogging(context.TODO(),
					&cloudtrail.StartLoggingInput{
						Name: trail.TrailARN,
					})
				if err != nil {
					log.Error().Err(err).Msg("failed to start trail")
				}
			}
		}

		thirdPartyReference := api.ThirdPartyReferenceResources{
			ServiceType: "AwsCloudTrailBucket",
			Path: api.ThirdPartyReferenceResourcesPath{
				Type:           "S3BucketPathExpression",
				BucketName:     s3BucketName,
				PathExpression: pathExpression,
			},
		}
		thirdPartyReference.Authentication = cloud.SetAWSAuthentication(awsId, awsKey, iamRoleArn)
		requestBodySchema.Source.ThirdPartyRef.Resources = append(requestBodySchema.Source.ThirdPartyRef.Resources, thirdPartyReference)
	} else {
		thirdPartyReference := api.ThirdPartyReferenceResources{
			ServiceType: "AwsCloudTrailBucket",
			Path: api.ThirdPartyReferenceResourcesPath{
				Type:           "S3BucketPathExpression",
				BucketName:     s3BucketName,
				PathExpression: pathExpression,
			},
		}
		thirdPartyReference.Authentication = cloud.SetAWSAuthentication(awsId, awsKey, iamRoleArn)
		requestBodySchema.Source.ThirdPartyRef.Resources = append(requestBodySchema.Source.ThirdPartyRef.Resources, thirdPartyReference)
	}
	// Send request to the HTTP endpoint to create CloudTrail source
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

	err = json.Unmarshal(responseBody, &awsCloudTrailResponse)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}

	if snsTopicArn != "" {
		snsClient := sns.NewFromConfig(cfg)
		snsClient.Subscribe(context.TODO(), &sns.SubscribeInput{
			Protocol: to.StringPtr("https"),
			TopicArn: to.StringPtr(snsTopicArn),
			Endpoint: to.StringPtr(awsCloudTrailResponse.Source.Url),
		})
	} else {
		log.Warn().Err(err).Msg("no snsTopicArn provided, sns message notification won't be enabled")
	}

	awsCloudTrailResponseJson, err := json.MarshalIndent(awsCloudTrailResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 201 {
		log.Error().Msg("failed to create source " + string(responseBody))
	} else {
		fmt.Println(string(awsCloudTrailResponseJson))
	}
}
