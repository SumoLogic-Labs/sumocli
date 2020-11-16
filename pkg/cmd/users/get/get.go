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

func NewCmdGetUser() *cobra.Command {
	var (
		id     string
		output string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic user",
		Long: `The following fields can be exported using the --output command:
firstName
lastName
email
roleIds
id
isActive
isLocked
isMfaEnabled
lastLoginTimestamp
`,
		Run: func(cmd *cobra.Command, args []string) {
			getUser(id, output)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user to get")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func getUser(id string, output string) {
	var userInfo api.UserResponse

	if id == "" {
		fmt.Println("--id field needs to be specified.")
		os.Exit(0)
	}

	requestUrl := "v1/users/" + id
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	util2.LogError(err)

	jsonErr := json.Unmarshal(responseBody, &userInfo)
	util2.LogError(jsonErr)

	userInfoJson, err := json.MarshalIndent(userInfo, "", "    ")
	util2.LogError(err)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody)
	} else {
		if factory.ValidateUserOutput(output) == true {
			value := gjson.Get(string(userInfoJson), output)
			formattedValue := strings.Trim(value.String(), `"[]"`)
			fmt.Println(formattedValue)
		} else {
			fmt.Println(string(userInfoJson))
		}
	}
}
