package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wizedkyle/sumocli/util"
	"io/ioutil"
	"net/http"
	"os"
)

type CreateUserRequest struct {
	Firstname    string   `json:"firstName"`
	Lastname     string   `json:"lastName"`
	Emailaddress string   `json:"email"`
	Roleids      []string `json:"roleIds"`
}

// TODO: complete this post creating a function to ListRoleIds
func CreateUser(firstName string, lastName string, emailAddress string, roleIds []string) {
	client := util.GetHttpClient()

	requestBodySchema := &CreateUserRequest{
		Firstname:    firstName,
		Lastname:     lastName,
		Emailaddress: emailAddress,
		Roleids:      roleIds,
	}

	requestBody, _ := json.Marshal(requestBodySchema)
	fmt.Println(string(requestBody))

	request, err := http.NewRequest("POST", util.GetApiEndpoint()+"v1/users", bytes.NewBuffer(requestBody))
	request.Header.Add("Authorization", util.GetApiCredentials())
	request.Header.Add("Content-Type", "application/json")
	util.LogError(err)

	response, err := client.Do(request)
	util.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	responseString := string(responseBody)
	fmt.Println(responseString)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	apiCallResult := util.HttpError(response.StatusCode, responseString)
	if apiCallResult == false {
		os.Exit(0)
	}
}
