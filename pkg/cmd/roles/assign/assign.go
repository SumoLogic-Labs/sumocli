package assign

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdRoleAssign(client *cip.APIClient) *cobra.Command {
	var (
		roleId string
		userId string
	)
	cmd := &cobra.Command{
		Use:   "assign",
		Short: "Assigns the specified Sumo Logic user to the role.",
		Run: func(cmd *cobra.Command, args []string) {
			assignRoleToUser(client, roleId, userId)
		},
	}
	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the identifier of the role to assign.")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the identifier of the user to assign the role to.")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func assignRoleToUser(client *cip.APIClient, roleId string, userId string) {
	apiResponse, httpResponse, errorResponse := client.AssignRoleToUser(roleId, userId)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
