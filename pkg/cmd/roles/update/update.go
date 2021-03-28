package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"reflect"
	"strconv"
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
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates a Sumo Logic role.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role update request started.")
			updateRole(id, name, description, filter, users, capabilities, autofill, merge)
			logger.Debug().Msg("Role update request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to update.")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name for the role.")
	cmd.Flags().StringVar(&description, "description", "", "Specify the role description.")
	cmd.Flags().StringVar(&filter, "filter", "", "Search filter for the role.")
	cmd.Flags().StringSliceVar(&users, "users", []string{}, "Comma deliminated list of user ids to add to the role.")
	cmd.Flags().StringSliceVar(&capabilities, "capabilities", []string{}, "Comma deliminated list of capabilities.")
	cmd.Flags().BoolVar(&autofill, "autofill", true, "Is set to true by default.")
	cmd.Flags().BoolVar(&merge, "merge", true, "Is set to true by default, if set to false it will overwrite the role.")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("filter")
	cmd.MarkFlagRequired("users")
	cmd.MarkFlagRequired("capabilities")
	return cmd
}

func updateRole(id string, name string, description string, filter string, users []string, capabilities []string, autofill bool, merge bool) {
	log := logging.GetConsoleLogger()
	var roleInfo api.RoleData
	if merge == true {
		requestUrl := "v1/roles/" + id
		client, request := factory.NewHttpRequest("GET", requestUrl)
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request " + requestUrl)
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}
		err = json.Unmarshal(responseBody, &roleInfo)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		if response.StatusCode != 200 {
			log.Fatal().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
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
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request " + requestUrl)
		}

		defer response.Body.Close()
		responseBody, err = io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}

		err = json.Unmarshal(responseBody, &roleInfo)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}

		roleInfoJson, err := json.MarshalIndent(roleInfo, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("error marshalling response body")
		}

		if response.StatusCode != 200 {
			log.Fatal().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
		} else {
			fmt.Println(string(roleInfoJson))
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
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request")
		}

		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}

		err = json.Unmarshal(responseBody, &roleInfo)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}

		roleInfoJson, err := json.MarshalIndent(roleInfo, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("error marshalling response body")
		}

		if response.StatusCode != 200 {
			log.Fatal().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
		} else {
			fmt.Println(string(roleInfoJson))
		}
	}
}
