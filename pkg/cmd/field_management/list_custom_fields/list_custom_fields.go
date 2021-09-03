package list_custom_fields

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdFieldManagementListCustomFields(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-custom-fields",
		Short: "Request a list of all the custom fields configured in your account.",
		Run: func(cmd *cobra.Command, args []string) {
			listCustomFields(client)
		},
	}
	return cmd
}

func listCustomFields(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.ListCustomFields()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
