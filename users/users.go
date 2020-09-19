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

type CreateUserResponse struct {
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Email              string   `json:"email"`
	RoleIds            []string `json:"roleIds"`
	CreatedAt          string   `json:"createdAt"`
	CreatedBy          string   `json:"createdBy"`
	ModifiedAt         string   `json:"modifiedAt"`
	ModifiedBy         string   `json:"modifiedBy"`
	Id                 string   `json:"id"`
	IsActive           bool     `json:"isActive"`
	IsLocked           bool     `json:"isLocked"`
	IsMfaEnabled       bool     `json:"isMfaEnabled"`
	LastLoginTimestamp string   `json:"lastLoginTimeStamp"`
}

func CreateUser(firstName string, lastName string, emailAddress string, roleIds []string) {
	var createUserResponse CreateUserResponse
	client := util.GetHttpClient()

	requestBodySchema := &CreateUserRequest{
		Firstname:    firstName,
		Lastname:     lastName,
		Emailaddress: emailAddress,
		Roleids:      roleIds,
	}

	requestBody, _ := json.Marshal(requestBodySchema)

	request, err := http.NewRequest("POST", util.GetApiEndpoint()+"v1/users", bytes.NewBuffer(requestBody))
	request.Header.Add("Authorization", util.GetApiCredentials())
	request.Header.Add("Content-Type", "application/json")
	util.LogError(err)

	response, err := client.Do(request)
	util.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	responseString := string(responseBody)

	apiCallResult := util.HttpError(response.StatusCode, responseString)
	if apiCallResult == false {
		os.Exit(0)
	} else if apiCallResult == true {
		jsonErr := json.Unmarshal(responseBody, &createUserResponse)
		util.LogError(jsonErr)
		fmt.Println("User account successfully created for " + createUserResponse.Firstname + " " + createUserResponse.Lastname)
	}
}
