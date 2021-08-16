package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
	apiResponse, httpResponse, errorResponse := client.GetDynamicParsingRule(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
