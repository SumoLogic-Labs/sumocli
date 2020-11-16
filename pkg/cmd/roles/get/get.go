package get

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
	"strings"
)

func NewCmdRoleGet() *cobra.Command {
	var (
		id     string
		output string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic role",
		Long: `The following fields can be exported using the --output command:
name
description
filterPredicate
users
capabilities
id
`,
		Run: func(cmd *cobra.Command, args []string) {
			getRole(id, output)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to get")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func getRole(id string, output string) {
	var roleInfo api.RoleData

	if id == "" {
		fmt.Println("--id field needs to be specified.")
		os.Exit(0)
	}

	requestUrl := "v1/roles/" + id
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	util2.LogError(err)

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
