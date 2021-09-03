package aws_cloudtrail

import (
	NewCmdAWSCloudTrailSourceCreate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/aws_cloudtrail/create"
	"github.com/spf13/cobra"
)

func NewCmdAWSCloudTrailSource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws-cloudtrail",
		Short: "Manage AWS CloudTrail sources",
	}
	cmd.AddCommand(NewCmdAWSCloudTrailSourceCreate.NewCmdAWSCloudTrailSourceCreate())
	return cmd
}
