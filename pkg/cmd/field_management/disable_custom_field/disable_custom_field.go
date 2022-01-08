package disable_custom_field

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdFieldManagementDisableCustomField(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use: "disable-custom-field",
		Short: "After disabling a field Sumo Logic will start dropping its incoming values at ingest. " +
			"As a result, they won't be searchable or usable. Historical values are not removed and remain searchable.",
		Run: func(cmd *cobra.Command, args []string) {
			disableCustomField(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field")
	cmd.MarkFlagRequired("id")
	return cmd
}

func disableCustomField(id string, client *cip.APIClient) {
	response, err := client.DisableField(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Field has been disabled.")
	}
}
