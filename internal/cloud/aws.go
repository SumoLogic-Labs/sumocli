package cloud

import (
	"github.com/SumoLogic-Labs/sumocli/api"
	"github.com/SumoLogic-Labs/sumocli/pkg/logging"
)

func SetAWSAuthentication(awsId string, awsKey string, roleArn string) api.ThirdPartyReferenceResourcesAuthentication {
	log := logging.GetConsoleLogger()
	authentication := api.ThirdPartyReferenceResourcesAuthentication{}
	if awsId != "" && awsKey != "" && roleArn == "" {
		authentication = api.ThirdPartyReferenceResourcesAuthentication{
			Type:   "S3BucketAuthentication",
			AwsId:  awsId,
			AwsKey: awsKey,
		}
	} else if awsId == "" && awsKey == "" && roleArn != "" {
		authentication = api.ThirdPartyReferenceResourcesAuthentication{
			Type:    "AWSRoleBasedAuthentication",
			RoleArn: roleArn,
		}
	} else {
		log.Fatal().Msg("Please enter either an AWS Id and AWS Key or an IAM Role ARN")
	}
	return authentication
}
