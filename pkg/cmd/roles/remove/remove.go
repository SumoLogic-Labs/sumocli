package remove

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/config"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
)

func NewCmdRoleRemoveUser() *cobra.Command {
	var (
		roleId string
		userId string
	)
	cmd := &cobra.Command{
		Use:   "remove user",
		Short: "Removes the specified Sumo Logic user from the role.",
		Run: func(cmd *cobra.Command, args []string) {
			removeUserRole(roleId, userId)
		},
	}
	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the identifier of the role")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the identifier of the user to remove")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func removeUserRole(roleId string, userId string) {
	client := config.GetSumoLogicSDKConfig()
	httpResponse, errorResponse := client.RemoveRoleFromUser(roleId, userId)
	if errorResponse != nil {
		fmt.Println(errorResponse.Error())
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "User: "+userId+" was removed from role: "+roleId)
	}
}
