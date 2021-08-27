package delete

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
