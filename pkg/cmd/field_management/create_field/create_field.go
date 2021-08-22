package create_field

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
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
