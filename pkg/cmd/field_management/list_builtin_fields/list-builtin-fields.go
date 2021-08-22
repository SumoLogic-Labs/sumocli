package list_builtin_fields

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdFieldManagementListBuiltinFields(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-builtin-fields",
		Short: "Built-in fields are created automatically by Sumo Logic for standard configuration purposes. " +
			"They include _sourceHost and _sourceCategory. Built-in fields can't be deleted or disabled.",
		Run: func(cmd *cobra.Command, args []string) {
			listBuiltinFields(client)
		},
	}
	return cmd
}

func listBuiltinFields(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.ListBuiltInFields()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
