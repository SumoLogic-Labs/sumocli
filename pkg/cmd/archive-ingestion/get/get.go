package get

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdArchiveIngestionGet(client *cip.APIClient) *cobra.Command {
	var (
		limit    int32
		sourceId string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a list of all the ingestion jobs created on an Archive Source.",
		Run: func(cmd *cobra.Command, args []string) {
			getArchiveIngestion(limit, sourceId, client)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 10, "Specify the number of jobs to return")
	cmd.Flags().StringVar(&sourceId, "sourceId", "", "Specify the id of the Archive Source")
	cmd.MarkFlagRequired("sourceId")
	return cmd
}

func getArchiveIngestion(limit int32, sourceId string, client *cip.APIClient) {
	var options types.ArchiveOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	data, response, err := client.ListArchiveJobsBySourceId(sourceId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = getArchiveIngestionPagination(client, options, paginationToken, sourceId)
		paginationToken = data.Next
	}
}

func getArchiveIngestionPagination(client *cip.APIClient, options types.ArchiveOpts, token string, sourceId string) types.ListArchiveJobsResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.ListArchiveJobsBySourceId(sourceId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
