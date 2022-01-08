package delete

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
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
	response, err := client.DeleteTable(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Lookup table deleted successfully.")
	}
}
