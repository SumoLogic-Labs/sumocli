package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func NewCmdRoleUpdate() *cobra.Command {
	var (
		id           string
		name         string
		description  string
		filter       string
		users        []string
		capabilities []string
		autofill     bool
		merge        bool
		output       string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			updateRole(id, name, description, filter, users, capabilities, autofill, merge, output)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to update")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name for the role")
	cmd.Flags().StringVar(&description, "description", "", "Specify the role description")
	cmd.Flags().StringVar(&filter, "filter", "", "Search filter for the role")
	cmd.Flags().StringSliceVar(&users, "users", []string{}, "Comma deliminated list of user ids to add to the role")
	cmd.Flags().StringSliceVar(&capabilities, "capabilities", []string{}, "Comma deliminated list of capabilities")
	cmd.Flags().BoolVar(&autofill, "autofill", true, "Is set to true by default.")
	cmd.Flags().BoolVar(&merge, "merge", true, "Is set to true by default, if set to false it will overwrite the role")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func updateRole(id string, name string, description string, filter string, users []string, capabilities []string, autofill bool, merge bool, output string) {
	var roleInfo api.RoleData
	if id == "" {
		fmt.Println("--id field needs to be set.")
		os.Exit(0)
	}

	if merge == true {
		requestUrl := "v1/roles/" + id
		client, request := factory.NewHttpRequest("GET", requestUrl)
		response, err := client.Do(request)
		util2.LogError(err)

		defer response.Body.Close()
		responseBody, err := ioutil.ReadAll(response.Body)
		util2.LogError(err)

		jsonErr := json.Unmarshal(responseBody, &roleInfo)
		util2.LogError(jsonErr)

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody)
			os.Exit(0)
		}

		// Building body payload to update the role based on the differences
		// between the current role settings and the desired settings
		requestBodySchema := &api.UpdateRoleRequest{}
		if strings.EqualFold(roleInfo.Name, name) {
			requestBodySchema.Name = roleInfo.Name
		} else {
			requestBodySchema.Name = name
		}

		if strings.EqualFold(roleInfo.Description, description) {
			requestBodySchema.Description = roleInfo.Description
		} else {
			requestBodySchema.Description = description
		}

		if strings.EqualFold(roleInfo.FilterPredicate, filter) {
			requestBodySchema.FilterPredicate = roleInfo.FilterPredicate
		} else {
			requestBodySchema.FilterPredicate = filter
		}

		if reflect.DeepEqual(roleInfo.Users, users) {
			requestBodySchema.Users = roleInfo.Users
		} else {
			requestBodySchema.Users = append(requestBodySchema.Users, roleInfo.Users...)
			requestBodySchema.Users = append(requestBodySchema.Users, users...)
		}

		if reflect.DeepEqual(roleInfo.Capabilities, capabilities) {
			requestBodySchema.Capabilities = roleInfo.Capabilities
		} else {
			requestBodySchema.Capabilities = append(requestBodySchema.Capabilities, roleInfo.Capabilities...)
			requestBodySchema.Capabilities = append(requestBodySchema.Capabilities, capabilities...)
		}

		if roleInfo.AutofillDependencies == autofill {
			requestBodySchema.AutoFillDependencies = roleInfo.AutofillDependencies
		} else {
			requestBodySchema.AutoFillDependencies = autofill
		}

		requestBody, _ := json.Marshal(requestBodySchema)
		client, request = factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
		response, err = client.Do(request)
		util2.LogError(err)

		defer response.Body.Close()
		responseBody, err = ioutil.ReadAll(response.Body)

		jsonErr = json.Unmarshal(responseBody, &roleInfo)
		util2.LogError(jsonErr)

		roleInfoJson, err := json.MarshalIndent(roleInfo, "", "    ")
		util2.LogError(err)

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody)
		} else {
			if factory.ValidateRoleOutput(output) == true {
				value := gjson.Get(string(roleInfoJson), output)
				formattedValue := strings.Trim(value.String(), `"[]"`)
				fmt.Println(formattedValue)
			} else {
				fmt.Println(string(roleInfoJson))
			}
		}
	} else {
		requestBodySchema := &api.UpdateRoleRequest{
			Name:                 name,
			Description:          description,
			FilterPredicate:      filter,
			Users:                users,
			Capabilities:         capabilities,
			AutoFillDependencies: autofill,
		}
		requestBody, _ := json.Marshal(requestBodySchema)

		requestUrl := "v1/roles/" + id
		client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
		response, err := client.Do(request)
		util2.LogError(err)

		defer response.Body.Close()
		responseBody, err := ioutil.ReadAll(response.Body)

		jsonErr := json.Unmarshal(responseBody, &roleInfo)
		util2.LogError(jsonErr)

		roleInfoJson, err := json.MarshalIndent(roleInfo, "", "    ")
		util2.LogError(err)

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody)
		} else {
			if factory.ValidateRoleOutput(output) == true {
				value := gjson.Get(string(roleInfoJson), output)
				formattedValue := strings.Trim(value.String(), `"[]"`)
				fmt.Println(formattedValue)
			} else {
				fmt.Println(string(roleInfoJson))
			}
		}
	}
}
