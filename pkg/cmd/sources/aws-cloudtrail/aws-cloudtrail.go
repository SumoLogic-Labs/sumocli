package aws_cloudtrail

import (
	"github.com/spf13/cobra"
)

func NewCmdAWSCloudTrailSource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws-cloudtrail",
		Short: "Manage AWS CloudTrail sources",
	}
	return cmd
}
