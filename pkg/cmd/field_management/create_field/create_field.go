package create_field

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdFieldManagementCreateField(client *cip.APIClient) *cobra.Command {
	var fieldName string
	cmd := &cobra.Command{
		Use:   "create-field",
		Short: "Adding a field will define it in the Fields schema allowing it to be assigned as metadata to your logs.",
		Run: func(cmd *cobra.Command, args []string) {
			createField(fieldName, client)
		},
	}
	cmd.Flags().StringVar(&fieldName, "fieldName", "", "Specify the name of the field")
	cmd.MarkFlagRequired("fieldName")
	return cmd
}

func createField(fieldName string, client *cip.APIClient) {
	data, response, err := client.CreateField(types.FieldName{
		FieldName: fieldName,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
