package get

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGetUser(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic user",
		Run: func(cmd *cobra.Command, args []string) {
			getUser(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user to get")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getUser(id string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetUser(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
