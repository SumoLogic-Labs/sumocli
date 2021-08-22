package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdFieldExtractionRulesGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a field extraction rule with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getFieldExtractionRule(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field extraction rule")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getFieldExtractionRule(id string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetExtractionRule(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
