package enable_custom_field

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdFieldManagementEnableCustomField(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use: "enable-custom-field",
		Short: "Fields have to be enabled to be assigned to your data. " +
			"This operation ensures that a specified field is enabled and Sumo Logic will treat it as safe to process. " +
			"All created custom fields using sumocli field_management create_field are enabled by default.",
		Run: func(cmd *cobra.Command, args []string) {
			enableCustomField(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field")
	cmd.MarkFlagRequired("id")
	return cmd
}

func enableCustomField(id string, client *cip.APIClient) {
	response, err := client.EnableField(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Field has been enabled.")
	}
}
