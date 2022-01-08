package delete

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdTransformationRulesDelete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a transformation rule with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			delete(client, id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the identifier of the transformation rule to delete.")
	cmd.MarkFlagRequired("id")
	return cmd
}

func delete(client *cip.APIClient, id string) {
	response, err := client.DeleteTransformationRule(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "The transformation rule was successfully deleted.")
	}
}
