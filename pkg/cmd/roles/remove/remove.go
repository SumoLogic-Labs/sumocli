package remove

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdRoleRemoveUser(client *cip.APIClient) *cobra.Command {
	var (
		roleId string
		userId string
	)
	cmd := &cobra.Command{
		Use:   "remove user",
		Short: "Removes the specified Sumo Logic user from the role.",
		Run: func(cmd *cobra.Command, args []string) {
			removeRoleFromUser(client, roleId, userId)
		},
	}
	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the identifier of the role.")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the identifier of the user to remove from the role.")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func removeRoleFromUser(client *cip.APIClient, roleId string, userId string) {
	httpResponse, errorResponse := client.RemoveRoleFromUser(roleId, userId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "User: "+userId+" was removed from role: "+roleId)
	}
}
