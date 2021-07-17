package archive_ingestion

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	NewCmdArchiveIngestionCreate "github.com/wizedkyle/sumocli/pkg/cmd/archive-ingestion/create"
	NewCmdArchiveIngestionDelete "github.com/wizedkyle/sumocli/pkg/cmd/archive-ingestion/delete"
	NewCmdArchiveIngestionGet "github.com/wizedkyle/sumocli/pkg/cmd/archive-ingestion/get"
	NewCmdArchiveIngestionList "github.com/wizedkyle/sumocli/pkg/cmd/archive-ingestion/list"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdArchiveIngestion(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "archive-ingestion",
		Short: "Manages archive ingestion",
		Long: "Archive Ingestion allows you to ingest data from Archive destinations. " +
			"You can use this command to ingest data from your Archive with an existing AWS S3 Archive Source.",
	}
	cmd.AddCommand(NewCmdArchiveIngestionCreate.NewCmdArchiveIngestionCreate(client, log))
	cmd.AddCommand(NewCmdArchiveIngestionDelete.NewCmdArchiveIngestionDelete())
	cmd.AddCommand(NewCmdArchiveIngestionGet.NewCmdArchiveIngestionGet(client, log))
	cmd.AddCommand(NewCmdArchiveIngestionList.NewCmdArchiveIngestionList(client, log))
	return cmd
}
