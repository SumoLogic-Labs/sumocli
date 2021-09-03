package list_dropped_fields

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdFieldManagementListDroppedFields(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list-dropped-fields",
		Short: "Dropped fields are fields sent to Sumo Logic, but are ignored since they are not defined in your Fields schema. " +
			"In order to save these values a field must both exist and be enabled.",
		Run: func(cmd *cobra.Command, args []string) {
			listDroppedFields(client)
		},
	}
	return cmd
}

func listDroppedFields(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.ListDroppedFields()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
