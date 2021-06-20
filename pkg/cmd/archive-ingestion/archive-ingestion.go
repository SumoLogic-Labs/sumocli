package archive_ingestion

import (
	"github.com/spf13/cobra"
	NewCmdArchiveIngestionList "github.com/wizedkyle/sumocli/pkg/cmd/archive-ingestion/list"
)

func NewCmdArchiveIngestion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "archive-ingestion",
		Short: "Manages archive ingestion",
		Long: "Archive Ingestion allows you to ingest data from Archive destinations. " +
			"You can use this command to ingest data from your Archive with an existing AWS S3 Archive Source.",
	}
	//cmd.AddCommand(NewCmdArchiveIngestionCreate.NewCmdArchiveIngestionCreate())
	//cmd.AddCommand(NewCmdArchiveIngestionGet.NewCmdArchiveIngestionGet())
	cmd.AddCommand(NewCmdArchiveIngestionList.NewCmdArchiveIngestionList())
	return cmd
}
