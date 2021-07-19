package aws_cloudtrail

import (
	"github.com/spf13/cobra"
	NewCmdAWSCloudTrailSourceCreate "github.com/wizedkyle/sumocli/pkg/cmd/sources/aws_cloudtrail/create"
)

func NewCmdAWSCloudTrailSource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws-cloudtrail",
		Short: "Manage AWS CloudTrail sources",
	}
	cmd.AddCommand(NewCmdAWSCloudTrailSourceCreate.NewCmdAWSCloudTrailSourceCreate())
	return cmd
}
