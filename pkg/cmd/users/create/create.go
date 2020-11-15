package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
)

func NewCmdUserCreate() *cobra.Command {
	var (
		firstName    string
		lastName     string
		emailAddress string
		roleIds      []string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic user account",
		Run: func(cmd *cobra.Command, args []string) {
			user(firstName, lastName, emailAddress, roleIds)
		},
	}

	cmd.Flags().StringVar(&firstName, "firstname", "", "First name of the user")
	cmd.Flags().StringVar(&lastName, "lastname", "", "Last name of the user")
	cmd.Flags().StringVar(&emailAddress, "email", "", "Email address of the user")
	cmd.Flags().StringSliceVar(&roleIds, "roleids", []string{}, "Comma deliminated list of Role Ids")

	return cmd
}

func user(firstName string, lastName string, emailAddress string, roleIds []string) {
	var createUserResponse api.CreateUserResponse

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
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody)
	} else {
		jsonErr := json.Unmarshal(responseBody, &createUserResponse)
		util2.LogError(jsonErr)
		fmt.Println("User account successfully created for " + createUserResponse.Firstname + " " + createUserResponse.Lastname)
	}
}
