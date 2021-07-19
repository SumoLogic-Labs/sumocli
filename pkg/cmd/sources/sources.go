package sources

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	cmdAwsCloudTrailSource "github.com/wizedkyle/sumocli/pkg/cmd/sources/aws_cloudtrail"
	cmdAWSS3ArchiveSource "github.com/wizedkyle/sumocli/pkg/cmd/sources/aws_s3_archive"
	cmdAzureEventHubSource "github.com/wizedkyle/sumocli/pkg/cmd/sources/azure_event_hub"
	cmdSourcesDelete "github.com/wizedkyle/sumocli/pkg/cmd/sources/delete"
	cmdHttpSources "github.com/wizedkyle/sumocli/pkg/cmd/sources/http"
	cmdSourcesList "github.com/wizedkyle/sumocli/pkg/cmd/sources/list"
	cmdLocalFileSources "github.com/wizedkyle/sumocli/pkg/cmd/sources/local-file"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdSources(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sources",
		Short: "Manages sources assigned to collectors",
	}
	cmd.AddCommand(cmdAwsCloudTrailSource.NewCmdAWSCloudTrailSource())
	cmd.AddCommand(cmdAWSS3ArchiveSource.NewCmdAWSS3ArchiveSource())
	cmd.AddCommand(cmdAzureEventHubSource.NewCmdAzureEventHubSource(client, log))
	cmd.AddCommand(cmdSourcesDelete.NewCmdDeleteSource())
	cmd.AddCommand(cmdHttpSources.NewCmdHttpSources())
	cmd.AddCommand(cmdSourcesList.NewCmdSourceList())
	cmd.AddCommand(cmdLocalFileSources.NewCmdLocalFileSources())
	return cmd
}
