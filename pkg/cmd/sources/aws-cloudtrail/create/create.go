package create

import (
	"context"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/internal/cloud"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdAWSCloudTrailSourceCreate() *cobra.Command {
	var (
		awsId                     string
		awsKey                    string
		createCloudTrailResources bool
		collectorId               int
		iamRoleArn                string
		name                      string
		pathExpression            string
		paused                    bool
		roleArn                   string
		s3BucketName              string
		snsTopicArn               string
		scanInterval              int
	)

	cmd := &cobra.Command{
		Use: "create",
		Short: "AWS CloudTrail records API calls made to AWS. " +
			"This includes calls made using the AWS Management Console, AWS SDKs, command line tools, and higher-level AWS services. " +
			"Add an AWS CloudTrail Source to upload these messages to Sumo Logic. " +
			"The AWS CloudTrail Source automatically parses the logs prior to upload.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.Flags().IntVar(&scanInterval, "scanInterval", 300000, "Specify the time interval of S3 bucket scans "+
		"for new data, in milliseconds. Minimum value is 1000. For automatic set value to -1.")
	return cmd
}

func createAWSCloudTrailSource(awsId string, awsKey string, createCloudTrailResources bool, collectorId int,
	iamRoleArn string, name string, pathExpression string, paused bool, roleArn string, s3BucketName string,
	snsTopicArn string, scanInterval int) {
	log := logging.GetConsoleLogger()
	var awsCloudTrailResponse api.AWSCloudTrail
	// TODO: Add filters and fields code here
	requestBodySchema := &api.AWSCloudTrail{
		SourceType:                 "Polling",
		Name:                       name,
		ContentType:                "AwsCloudTrailBucket",
		ScanInterval:               scanInterval,
		Paused:                     paused,
		AutomaticDateParsing:       false,
		MultilineProcessingEnabled: false,
		UseAutolineMatching:        false,
		ForceTimeZone:              false,
		Filters:                    nil,
		CutoffTimestamp:            0,
		Encoding:                   "",
		Fields:                     nil,
	}

	if createCloudTrailResources == true {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal().Err(err).Msg("failed to load AWS config")
		}

		s3Client := s3.Client{}
		bucket, err := s3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: nil,
		})

		cloudtrailClient := cloudtrail.Client{}
		trail, err := cloudtrailClient.CreateTrail(
			context.TODO(), &cloudtrail.CreateTrailInput{
				Name:                       nil,
				S3BucketName:               nil,
				CloudWatchLogsLogGroupArn:  nil,
				CloudWatchLogsRoleArn:      nil,
				EnableLogFileValidation:    nil,
				IncludeGlobalServiceEvents: nil,
				IsMultiRegionTrail:         to.BoolPtr(true),
				IsOrganizationTrail:        nil,
				KmsKeyId:                   nil,
				S3KeyPrefix:                nil,
				SnsTopicName:               nil,
				TagsList:                   nil,
			})
		if err != nil {
			log.Error().Err(err).Msg("failed to create AWS CloudTrail")
		}

		thirdPartyReference := api.ThirdPartyReferenceResources{
			ServiceType: "",
			Path: api.ThirdPartyReferenceResourcesPath{
				Type:           "",
				BucketName:     "",
				PathExpression: "",
				SnsTopicOrSubscriptionArn: api.ThirdPartyReferenceResourcesPathSnsTopic{
					IsSuccess: false,
					Arn:       "",
				},
			},
		}
		thirdPartyReference.Authentication = cloud.SetAWSAuthentication(awsId, awsKey, roleArn)
		requestBodySchema.ThirdPartyRef.Resources = append(requestBodySchema.ThirdPartyRef.Resources, thirdPartyReference)
	} else {
		thirdPartyReference := api.ThirdPartyReferenceResources{
			ServiceType: "",
			Path: api.ThirdPartyReferenceResourcesPath{
				Type:           "S3BucketPathExpression",
				BucketName:     s3BucketName,
				PathExpression: pathExpression,
				SnsTopicOrSubscriptionArn: api.ThirdPartyReferenceResourcesPathSnsTopic{
					IsSuccess: false,
					Arn:       snsTopicArn,
				},
			},
		}
		thirdPartyReference.Authentication = cloud.SetAWSAuthentication(awsId, awsKey, roleArn)
		requestBodySchema.ThirdPartyRef.Resources = append(requestBodySchema.ThirdPartyRef.Resources, thirdPartyReference)
	}

}
