package delete_data

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdLookupTablesDeleteData(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete-data",
		Short: "Delete all data from a lookup table.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteLookupTableData(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table to delete data from")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteLookupTableData(id string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.TruncateTable(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
