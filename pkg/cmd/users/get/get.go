package get

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
