package remove

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdRoleRemoveUser(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var (
		roleId string
		userId string
	)
	cmd := &cobra.Command{
		Use:   "remove user",
		Short: "Removes the specified Sumo Logic user from the role.",
		Run: func(cmd *cobra.Command, args []string) {
			removeRoleFromUser(client, roleId, userId, log)
		},
	}
	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the identifier of the role.")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the identifier of the user to remove from the role.")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func removeRoleFromUser(client *cip.APIClient, roleId string, userId string, log *zerolog.Logger) {
	httpResponse, errorResponse := client.RemoveRoleFromUser(roleId, userId)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to remove role from user")
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "User: "+userId+" was removed from role: "+roleId)
	}
}
