package create

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
	"strings"
)

func NewCmdUserCreate() *cobra.Command {
	var (
		firstName    string
		lastName     string
		emailAddress string
		roleIds      []string
		output       string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic user account",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User create request started.")
			user(firstName, lastName, emailAddress, roleIds, output, logger)
			logger.Debug().Msg("User create request finished.")
		},
	}

	cmd.Flags().StringVar(&firstName, "firstname", "", "First name of the user")
	cmd.Flags().StringVar(&lastName, "lastname", "", "Last name of the user")
	cmd.Flags().StringVar(&emailAddress, "email", "", "Email address of the user")
	cmd.Flags().StringSliceVar(&roleIds, "roleids", []string{}, "Comma deliminated list of Role Ids")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func user(firstName string, lastName string, emailAddress string, roleIds []string, output string, logger zerolog.Logger) {
	var createUserResponse api.UserResponse

	requestBodySchema := &api.CreateUserRequest{
		Firstname:    firstName,
		Lastname:     lastName,
		Emailaddress: emailAddress,
		Roleids:      roleIds,
	}
	requestBody, _ := json.Marshal(requestBodySchema)
	fmt.Println(string(requestBody))
	client, request := factory.NewHttpRequestWithBody("POST", "v1/users", requestBody)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	jsonErr := json.Unmarshal(responseBody, &createUserResponse)
	logging.LogError(jsonErr, logger)

	createUserResponseJson, err := json.MarshalIndent(createUserResponse, "", "    ")
	logging.LogError(err, logger)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		if factory.ValidateUserOutput(output) == true {
			value := gjson.Get(string(createUserResponseJson), output)
			formattedValue := strings.Trim(value.String(), `"[]"`)
			fmt.Println(formattedValue)
		} else {
			fmt.Println(string(createUserResponseJson))
			fmt.Println("User account successfully created for " + createUserResponse.Firstname + " " + createUserResponse.Lastname)
		}
	}
}
