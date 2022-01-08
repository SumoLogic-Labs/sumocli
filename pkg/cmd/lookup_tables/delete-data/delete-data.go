package delete_data

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
	data, response, err := client.TruncateTable(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
