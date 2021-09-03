package delete

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdFieldExtractionRulesDelete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a field extraction rule with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteFieldExtractionRules(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field extraction rule")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteFieldExtractionRules(id string, client *cip.APIClient) {
	httpResponse, errorResponse := client.DeleteExtractionRule(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Extraction rule was deleted successfully.")
	}
}
