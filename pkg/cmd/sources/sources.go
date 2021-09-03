package sources

import (
	cmdAwsCloudTrailSource "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/aws_cloudtrail"
	cmdAWSS3ArchiveSource "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/aws_s3_archive"
	cmdAzureEventHubSource "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/azure_event_hub"
	cmdSourcesDelete "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/delete"
	cmdHttpSources "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/http"
	cmdSourcesList "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/list"
	cmdLocalFileSources "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/sources/local-file"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdSources(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sources",
		Short: "Manages sources assigned to collectors",
	}
	cmd.AddCommand(cmdAwsCloudTrailSource.NewCmdAWSCloudTrailSource())
	cmd.AddCommand(cmdAWSS3ArchiveSource.NewCmdAWSS3ArchiveSource())
	cmd.AddCommand(cmdAzureEventHubSource.NewCmdAzureEventHubSource(client))
	cmd.AddCommand(cmdSourcesDelete.NewCmdDeleteSource())
	cmd.AddCommand(cmdHttpSources.NewCmdHttpSources())
	cmd.AddCommand(cmdSourcesList.NewCmdSourceList())
	cmd.AddCommand(cmdLocalFileSources.NewCmdLocalFileSources())
	return cmd
}
