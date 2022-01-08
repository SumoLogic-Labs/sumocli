package get

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdDynamicParsingGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a dynamic parsing rule with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getDynamicParsingRules(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the dynamic parsing rule")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getDynamicParsingRules(id string, client *cip.APIClient) {
	data, response, err := client.GetDynamicParsingRule(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
