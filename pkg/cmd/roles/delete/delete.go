package delete

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

func NewCmdRoleDelete() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			deleteRole(id)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to delete")

	return cmd
}

func deleteRole(id string) {
	if id == "" {
		fmt.Println("--id field needs to be set.")
		os.Exit(0)
	}

	requestUrl := "v1/roles/" + id
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
		if responseError.Errors[0].Code == "acl:role_has_users" {
			fmt.Println("The role wasn't deleted as users are assigned to it." +
				" Please run sumocli roles remove user then re-run sumocli roles delete")
		}
	} else {
		fmt.Println("Role was deleted.")
	}
}
