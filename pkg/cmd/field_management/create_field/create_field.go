package create_field

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
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
	apiResponse, httpResponse, errorResponse := client.CreateField(types.FieldName{
		FieldName: fieldName,
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
