package get_builtin_field

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
	apiResponse, httpResponse, errorResponse := client.GetBuiltInField(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
