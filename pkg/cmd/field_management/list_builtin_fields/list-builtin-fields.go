package list_builtin_fields

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
	data, response, err := client.ListBuiltInFields()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
