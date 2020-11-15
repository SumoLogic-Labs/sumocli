package remove

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
	"os"
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

	cmd.Flags().StringVar(&roleId, "roleid", "", "Specify the id of the role")
	cmd.Flags().StringVar(&userId, "userid", "", "Specify the id of the user to remove")

	return cmd
}

func removeUserRole(roleId string, userId string) {
	if roleId == "" || userId == "" {
		fmt.Println("--roleid and --userid fields need to be set")
		os.Exit(0)
	}

	requestUrl := "v1/roles/" + roleId + "/users/" + userId
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	util2.LogError(err)

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		jsonErr := json.Unmarshal(responseBody, &responseError)
		util2.LogError(jsonErr)
		fmt.Println(responseError.Errors[0].Message)
	} else {
		fmt.Println("User: " + userId + " was removed from role: " + roleId)
	}
}
