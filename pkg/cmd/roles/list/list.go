package list

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
	"github.com/wizedkyle/sumocli/pkg/cmd/roles"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
	"net/http"
	"net/url"
)

func NewCmdRoleList() *cobra.Command {
	var (
		numberOfResults string
		filter          string
		output          bool
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic roles",
		Run: func(cmd *cobra.Command, args []string) {
			listRoles(numberOfResults, filter, output)
		},
	}

	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")
	cmd.Flags().BoolVar(&output, "output", false, "Output results to a file, defaults to false")

	return cmd
}

func listRoles(numberOfResults string, name string, output bool) {
	var roleInfo roles.Role
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

	// Determines if the response should be written to a file or to console
	if output == true {
		util2.OutputToFile(roleInfoJson)
	} else {
		fmt.Println(string(roleInfoJson))
	}
}
