package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdLookupTablesGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic lookup table based on the given identifier",
		Run: func(cmd *cobra.Command, args []string) {
			getLookupTable(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table you want to retrieve")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getLookupTable(id string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.LookupTableById(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
