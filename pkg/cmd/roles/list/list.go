package list

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/config"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdRoleList() *cobra.Command {
	var (
		limit  int32
		name   string
		sortBy bool
	)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic roles",
		Run: func(cmd *cobra.Command, args []string) {
			listRoles(limit, name, sortBy)
		},
	}
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the role you want to retrieve")
	cmd.Flags().BoolVar(&sortBy, "sortBy", false, "Sorts the roles by the name field")
	return cmd
}

func listRoles(limit int32, name string, sortBy bool) {
	var sortByField string
	if sortBy == true {
		sortByField = "name"
	}
	client := config.GetSumoLogicSDKConfig()
	apiResponse, httpResponse, errorResponse := client.ListRoles(
		&types.ListRolesOpts{
			Limit:  optional.NewInt32(limit),
			SortBy: optional.NewString(sortByField),
			Name:   optional.NewString(name),
		})

	if errorResponse != nil {
		fmt.Println(errorResponse.Error())
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse)
	}
}

func listRolesRequest(token string) {

}
