package delete

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
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
	response, err := client.DeleteRole(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Role was deleted successfully")
	}
}
