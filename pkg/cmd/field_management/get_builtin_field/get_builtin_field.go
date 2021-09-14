package get_builtin_field

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdFieldManagementGetBuiltinField(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get-builtin-field",
		Short: "Get the details of a built-in field.",
		Run: func(cmd *cobra.Command, args []string) {
			getBuiltinField(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the builtin field")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getBuiltinField(id string, client *cip.APIClient) {
	data, response, err := client.GetBuiltInField(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
