package list

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io/ioutil"
	"net/url"
	"strings"
)

func NewCmdUserList() *cobra.Command {
	var (
		email           string
		numberOfResults string
		sortBy          string
		output          string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic users",
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
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User list request started.")
			listUsers(email, numberOfResults, sortBy, output, logger)
			logger.Debug().Msg("User list request finished.")
		},
	}

	cmd.Flags().StringVar(&email, "email", "", "Specify the email address of the user")
	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&sortBy, "sort", "", "Sort the results by firstName, lastName or email")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func listUsers(email string, numberOfResults string, sortBy string, output string, logger zerolog.Logger) {
	var userInfo api.Users

	client, request := factory.NewHttpRequest("GET", "v1/users")
	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if sortBy != "" {
		if factory.ValidateUserSortBy(sortBy) == false {
			fmt.Println(sortBy + "is an invalid field to sort by. Available fields are firstName, lastName or email. ")
		} else {
			query.Add("sortBy", sortBy)
		}
	}
	if email != "" {
		query.Add("email", email)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	logging.LogError(err, logger)

	jsonErr := json.Unmarshal(responseBody, &userInfo)
	logging.LogError(jsonErr, logger)

	userInfoJson, err := json.MarshalIndent(userInfo.Data, "", "    ")
	logging.LogError(err, logger)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		if factory.ValidateUserOutput(output) == true {
			value := gjson.Get(string(userInfoJson), "#."+output)
			formattedValue := strings.Trim(value.String(), `[]`)
			formattedValue = strings.ReplaceAll(formattedValue, "\"", "")
			fmt.Println(formattedValue)
		} else {
			fmt.Println(string(userInfoJson))
		}
	}
}
