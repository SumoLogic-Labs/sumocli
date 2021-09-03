package get_custom_field

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdFieldManagementGetCustomField(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get-custom-field",
		Short: "Get the details of a custom field.",
		Run: func(cmd *cobra.Command, args []string) {
			getCustomField(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the custom field")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getCustomField(id string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetCustomField(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
