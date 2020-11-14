package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
	"net/http"
	"os"
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
	cmd.Flags().StringArrayVar(&roleIds, "roleids", []string{}, "Comma deliminated list of Role Ids")

	return cmd
}

func user(firstName string, lastName string, emailAddress string, roleIds []string) {
	var createUserResponse api.CreateUserResponse
	client := util2.GetHttpClient()
	authToken, apiEndpoint := login.ReadCredentials()

	requestBodySchema := &api.CreateUserRequest{
		Firstname:    firstName,
		Lastname:     lastName,
		Emailaddress: emailAddress,
		Roleids:      roleIds,
	}

	requestBody, _ := json.Marshal(requestBodySchema)

	request, err := http.NewRequest("POST", apiEndpoint+"v1/users", bytes.NewBuffer(requestBody))
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	util2.LogError(err)

	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	apiCallResult := factory.HttpError(response.StatusCode)
	if apiCallResult == false {
		os.Exit(0)
	} else if apiCallResult == true {
		jsonErr := json.Unmarshal(responseBody, &createUserResponse)
		util2.LogError(jsonErr)
		fmt.Println("User account successfully created for " + createUserResponse.Firstname + " " + createUserResponse.Lastname)
	}
}
