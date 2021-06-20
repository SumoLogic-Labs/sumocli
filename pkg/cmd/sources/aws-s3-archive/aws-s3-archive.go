package aws_s3_archive

import (
	"github.com/spf13/cobra"
	NewCmdAWSS3ArchiveSourceCreate "github.com/wizedkyle/sumocli/pkg/cmd/sources/aws-s3-archive/create"
)

func NewCmdAWSS3ArchiveSource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws-s3-archive",
		Short: "Manage AWS S3 Archive sources",
	}
	cmd.AddCommand(NewCmdAWSS3ArchiveSourceCreate.NewCmdAWSS3ArchiveSourceCreate())
	return cmd
}
