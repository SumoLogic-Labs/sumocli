package list_custom_fields

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
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
	data, response, err := client.ListCustomFields()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
