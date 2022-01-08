package delete

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdDynamicParsingDelete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a dynamic parsing rule with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteDynamicParsingRules(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the dynamic parsing rule")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteDynamicParsingRules(id string, client *cip.APIClient) {
	response, err := client.DeleteDynamicParsingRule(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Dynamic Parsing rule was deleted successfully.")
	}
}
