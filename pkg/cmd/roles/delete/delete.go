package delete

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdRoleDelete(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			deleteRole(client, id, log)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the identifier of the role to delete.")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteRole(client *cip.APIClient, id string, log *zerolog.Logger) {
	httpResponse, errorResponse := client.DeleteRole(id)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to delete role")
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Role was deleted successfully")
	}
}
