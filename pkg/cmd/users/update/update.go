package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
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
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			updateUser(id, firstName, lastName, isActive, roleIds, merge)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user to update.")
	cmd.Flags().StringVar(&firstName, "firstname", "", "First name of the user.")
	cmd.Flags().StringVar(&lastName, "lastname", "", "Last name for the user.")
	cmd.Flags().BoolVar(&isActive, "isActive", true, "True if the account is active, false if it is deactivated")
	cmd.Flags().StringSliceVar(&roleIds, "roleIds", []string{}, "Comma deliminated list of Role Ids.")
	cmd.Flags().BoolVar(&merge, "merge", true, "Is set to true by default, if set to false it will overwrite the role.")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("firstname")
	cmd.MarkFlagRequired("lastname")
	cmd.MarkFlagRequired("isactive")
	cmd.MarkFlagRequired("roleids")
	return cmd
}

func updateUser(id string, firstName string, lastName string, isActive bool, roleIds []string, merge bool) {
	var userInfo api.UserResponse
	log := logging.GetConsoleLogger()
	if merge == true {
		requestUrl := "v1/users/" + id
		client, request := factory.NewHttpRequest("GET", requestUrl)
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request")
		}

		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}

		err = json.Unmarshal(responseBody, &userInfo)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, log)
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
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request")
		}

		defer response.Body.Close()
		responseBody, err = io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}

		err = json.Unmarshal(responseBody, &userInfo)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}

		userInfoJson, err := json.MarshalIndent(userInfo, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal response")
		}

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, log)
		} else {
			fmt.Println(string(userInfoJson))
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
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request")
		}

		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}

		err = json.Unmarshal(responseBody, &userInfo)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}

		userInfoJson, err := json.MarshalIndent(userInfo, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal response")
		}

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, log)
		} else {
			fmt.Println(string(userInfoJson))
		}
	}
}
