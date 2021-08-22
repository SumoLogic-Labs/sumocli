package delete_field

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdFieldManagementDeleteField(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use: "delete-field",
		Short: "Deleting a field does not delete historical data assigned with that field. " +
			"If you delete a field by mistake and one or more of those dependencies break, you can re-add the field to get things working properly again. " +
			"You should always disable a field using sumocli field_management disable_custom_field and ensure things are behaving as expected before deleting a field.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteField(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteField(id string, client *cip.APIClient) {
	httpResponse, errorResponse := client.DeleteField(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "The field was successfully deleted.")
	}
}
