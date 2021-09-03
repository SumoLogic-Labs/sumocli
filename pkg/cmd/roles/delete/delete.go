package delete

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdRoleDelete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			deleteRole(client, id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the identifier of the role to delete.")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteRole(client *cip.APIClient, id string) {
	httpResponse, errorResponse := client.DeleteRole(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Role was deleted successfully")
	}
}
