package aws_s3_archive

import (
	NewCmdAWSS3ArchiveSourceCreate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/sources/aws_s3_archive/create"
	"github.com/spf13/cobra"
)

func NewCmdAWSS3ArchiveSource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws-s3-archive",
		Short: "Manage AWS S3 Archive sources",
	}
	cmd.AddCommand(NewCmdAWSS3ArchiveSourceCreate.NewCmdAWSS3ArchiveSourceCreate())
	return cmd
}
