package delete

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
	httpResponse, errorResponse := client.DeleteDynamicParsingRule(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Dynamic Parsing rule was deleted successfully.")
	}
}
