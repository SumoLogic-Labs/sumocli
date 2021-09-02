package list_custom_fields

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
