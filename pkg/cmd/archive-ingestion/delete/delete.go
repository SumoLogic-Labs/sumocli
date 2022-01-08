package delete

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdArchiveIngestionDelete(client *cip.APIClient) *cobra.Command {
	var (
		id       string
		sourceId string
	)
	cmd := &cobra.Command{
		Use: "delete",
		Short: "Delete an ingestion job with the given identifier from the organization. " +
			"The delete operation is only possible for jobs with a Succeeded or Failed status.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteArchiveIngestion(id, sourceId, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the archive source")
	cmd.Flags().StringVar(&sourceId, "sourceId", "", "Specify the source Id of the Archive Source")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("sourceId")
	return cmd
}

func deleteArchiveIngestion(id string, sourceId string, client *cip.APIClient) {
	response, err := client.DeleteArchiveJob(sourceId, id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "The ingestion job was deleted successfully.")
	}
}
