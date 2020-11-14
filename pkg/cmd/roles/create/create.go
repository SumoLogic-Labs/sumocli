package create

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
	client, request := factory.NewHttpRequestWithBody("POST", "v1/roles", requestBody)
	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	apiCallResult := factory.HttpError(response.StatusCode)
	if apiCallResult == false {
		os.Exit(0)
	} else if apiCallResult == true {
		jsonErr := json.Unmarshal(responseBody, &createRoleResponse)
		util2.LogError(jsonErr)
		fmt.Println(createRoleResponse.Name + " role successfully created")
	}
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
