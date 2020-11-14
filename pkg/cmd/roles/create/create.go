package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
	"net/http"
)

func NewCmdRoleCreate() *cobra.Command {
	var (
		name         string
		description  string
		filter       string
		users        []string
		capabilities []string
		autofill     bool
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			createRole(name, description, filter, users, capabilities, autofill)
		},
	}

	return cmd
}

func createRole(name string, description string, filter string, users []string, capabilities []string, autofill bool) {
	var createRoleResponse api.RoleData
	client := util2.GetHttpClient()
	authToken, apiEndpoint := login.ReadCredentials()

	for i, capability := range capabilities {
		if validateCapabilities(capability) == false {
			fmt.Println(capability + " is not a valid Sumo Logic role capability.")
		}
		i++
	}

	requestBodySchema := &api.CreateRoleRequest{
		Name:                 name,
		Description:          description,
		FilterPredicate:      filter,
		Users:                users,
		Capabilities:         capabilities,
		AutoFillDependencies: autofill,
	}
	requestBody, _ := json.Marshal(requestBodySchema)
	request, err := http.NewRequest("POST", apiEndpoint+"v1/roles", bytes.NewBuffer(requestBody))
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	util2.LogError(err)

	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	responseString := string(responseBody)

}

func validateCapabilities(capability string) bool {
	switch capability {
	case
		"test",
		"test2":
		return true
	}
	return false
}
