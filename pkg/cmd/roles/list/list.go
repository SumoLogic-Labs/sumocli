package list

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func NewCmdRoleList() *cobra.Command {
	var (
		numberOfResults string
		filter          string
		output          string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic roles",
		Long: `The following fields can be exported using the --output command:
name
description
filterPredicate
users
capabilities
id
`,
		Run: func(cmd *cobra.Command, args []string) {
			listRoles(numberOfResults, filter, output)
		},
	}

	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func listRoles(numberOfResults string, name string, output string) {
	var roleInfo api.Role
	client := util2.GetHttpClient()
	authToken, apiEndpoint := login.ReadCredentials()

	request, err := http.NewRequest("GET", apiEndpoint+"v1/roles", nil)
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	util2.LogError(err)

	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if name != "" {
		query.Add("name", name)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	util2.LogError(err)

	jsonErr := json.Unmarshal(responseBody, &roleInfo)
	util2.LogError(jsonErr)

	roleInfoJson, err := json.MarshalIndent(roleInfo.Data, "", "    ")
	util2.LogError(err)

	if validateOutput(output) == true {
		value := gjson.Get(string(roleInfoJson), "#."+output)
		formattedValue := strings.Trim(value.String(), `"[]"`)
		fmt.Println(formattedValue)
	} else {
		fmt.Println(string(roleInfoJson))
	}
}

func validateOutput(output string) bool {
	switch output {
	case
		"name",
		"description",
		"filterPredicate",
		"users",
		"capabilities",
		"id":
		return true
	}
	return false
}
