package assign

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/config"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
)

func NewCmdRoleAssign() *cobra.Command {
	var (
		roleId string
		userId string
	)
	cmd := &cobra.Command{
		Use:   "assign",
		Short: "Assigns the specified Sumo Logic user to the role.",
		Run: func(cmd *cobra.Command, args []string) {
			assignRole(roleId, userId)
		},
	}
	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the identifier of the role to assign.")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the identifier of the user to assign the role to.")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func assignRole(roleId string, userId string) {
	client := config.GetSumoLogicSDKConfig()
	apiResponse, httpResponse, errorResponse := client.AssignRoleToUser(roleId, userId)
	if errorResponse != nil {
		fmt.Println(errorResponse.Error())
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
