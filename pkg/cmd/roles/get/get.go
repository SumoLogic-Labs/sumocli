package get

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdRoleGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic role information",
		Run: func(cmd *cobra.Command, args []string) {
			getRole(client, id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to get")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getRole(client *cip.APIClient, id string) {
	apiResponse, httpResponse, errorResponse := client.GetRole(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
