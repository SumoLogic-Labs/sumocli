package delete

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdLookupTablesDelete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a lookup table completely.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteLookupTable(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table to delete")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteLookupTable(id string, client *cip.APIClient) {
	httpResponse, errorResponse := client.DeleteTable(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Lookup table deleted successfully.")
	}
}
