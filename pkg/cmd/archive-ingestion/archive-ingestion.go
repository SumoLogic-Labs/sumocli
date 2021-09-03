package archive_ingestion

import (
	NewCmdArchiveIngestionCreate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/archive-ingestion/create"
	NewCmdArchiveIngestionDelete "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/archive-ingestion/delete"
	NewCmdArchiveIngestionGet "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/archive-ingestion/get"
	NewCmdArchiveIngestionList "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/archive-ingestion/list"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdArchiveIngestion(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "archive-ingestion",
		Short: "Manages archive ingestion",
		Long: "Archive Ingestion allows you to ingest data from Archive destinations. " +
			"You can use this command to ingest data from your Archive with an existing AWS S3 Archive Source.",
	}
	cmd.AddCommand(NewCmdArchiveIngestionCreate.NewCmdArchiveIngestionCreate(client))
	cmd.AddCommand(NewCmdArchiveIngestionDelete.NewCmdArchiveIngestionDelete(client))
	cmd.AddCommand(NewCmdArchiveIngestionGet.NewCmdArchiveIngestionGet(client))
	cmd.AddCommand(NewCmdArchiveIngestionList.NewCmdArchiveIngestionList(client))
	return cmd
}
