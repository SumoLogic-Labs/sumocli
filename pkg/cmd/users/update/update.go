package update

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
	"os"
	"reflect"
	"strings"
)

func NewCmdUserUpdate() *cobra.Command {
	var (
		id        string
		firstName string
		lastName  string
		isActive  bool
		roleIds   []string
		merge     bool
		output    string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Update user request started.")
			updateUser(id, firstName, lastName, isActive, roleIds, merge, output, logger)
			logger.Debug().Msg("Update user request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user to update.")
	cmd.Flags().StringVar(&firstName, "firstname", "", "First name of the user.")
	cmd.Flags().StringVar(&lastName, "lastname", "", "Last name for the user.")
	cmd.Flags().BoolVar(&isActive, "isactive", true, "True if the account is active, false if it is deactivated")
	cmd.Flags().StringSliceVar(&roleIds, "roleids", []string{}, "Comma deliminated list of Role Ids.")
	cmd.Flags().BoolVar(&merge, "merge", true, "Is set to true by default, if set to false it will overwrite the role.")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from.")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("firstname")
	cmd.MarkFlagRequired("lastname")
	cmd.MarkFlagRequired("isactive")
	cmd.MarkFlagRequired("roleids")
	return cmd
}

func updateUser(id string, firstName string, lastName string, isActive bool, roleIds []string, merge bool, output string, logger zerolog.Logger) {
	var userInfo api.UserResponse

	if merge == true {
		requestUrl := "v1/users/" + id
		client, request := factory.NewHttpRequest("GET", requestUrl)
		response, err := client.Do(request)
		logging.LogError(err, logger)

		defer response.Body.Close()
		responseBody, err := ioutil.ReadAll(response.Body)
		logging.LogError(err, logger)

		jsonErr := json.Unmarshal(responseBody, &userInfo)
		logging.LogError(jsonErr, logger)

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, logger)
			os.Exit(0)
		}

		// Building body payload to update the user based on the differences
		// between the current role settings and the desired settings
		requestBodySchema := &api.UpdateUserRequest{}
		requestBodySchema.IsActive = userInfo.IsActive
		if strings.EqualFold(userInfo.Firstname, firstName) {
			requestBodySchema.Firstname = userInfo.Firstname
		} else {
			requestBodySchema.Firstname = firstName
		}

		if strings.EqualFold(userInfo.Lastname, lastName) {
			requestBodySchema.Lastname = userInfo.Lastname
		} else {
			requestBodySchema.Lastname = lastName
		}

		if reflect.DeepEqual(userInfo.RoleIds, roleIds) {
			requestBodySchema.Roleids = userInfo.RoleIds
		} else {
			requestBodySchema.Roleids = append(requestBodySchema.Roleids, userInfo.RoleIds...)
			requestBodySchema.Roleids = append(requestBodySchema.Roleids, roleIds...)
		}

		requestBody, _ := json.Marshal(requestBodySchema)
		client, request = factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
		response, err = client.Do(request)
		logging.LogError(err, logger)

		defer response.Body.Close()
		responseBody, err = ioutil.ReadAll(response.Body)

		jsonErr = json.Unmarshal(responseBody, &userInfo)
		logging.LogError(jsonErr, logger)

		userInfoJson, err := json.MarshalIndent(userInfo, "", "    ")
		logging.LogError(jsonErr, logger)

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, logger)
		} else {
			if factory.ValidateUserOutput(output) == true {
				value := gjson.Get(string(userInfoJson), output)
				formattedValue := strings.Trim(value.String(), `"[]"`)
				fmt.Println(formattedValue)
			} else {
				fmt.Println(string(userInfoJson))
			}
		}
	} else {
		requestBodySchema := &api.UpdateUserRequest{
			Firstname: firstName,
			Lastname:  lastName,
			IsActive:  isActive,
			Roleids:   roleIds,
		}
		requestBody, _ := json.Marshal(requestBodySchema)

		requestUrl := "v1/users/" + id
		client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
		response, err := client.Do(request)
		logging.LogError(err, logger)

		defer response.Body.Close()
		responseBody, err := ioutil.ReadAll(response.Body)

		jsonErr := json.Unmarshal(responseBody, &userInfo)
		logging.LogError(jsonErr, logger)

		userInfoJson, err := json.MarshalIndent(userInfo, "", "    ")
		logging.LogError(err, logger)

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, logger)
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
}
