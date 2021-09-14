package get

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgetsGet(client *cip.APIClient) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get an ingest budget by the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getIngestBudget(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getIngestBudget(id string, client *cip.APIClient) {
	data, response, err := client.GetIngestBudget(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
