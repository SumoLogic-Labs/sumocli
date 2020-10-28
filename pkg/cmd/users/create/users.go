package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
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
	client := util2.GetHttpClient()

	requestBodySchema := &CreateUserRequest{
		Firstname:    firstName,
		Lastname:     lastName,
		Emailaddress: emailAddress,
		Roleids:      roleIds,
	}

	requestBody, _ := json.Marshal(requestBodySchema)

	request, err := http.NewRequest("POST", util2.GetApiEndpoint()+"v1/users", bytes.NewBuffer(requestBody))
	request.Header.Add("Authorization", util2.GetApiCredentials())
	request.Header.Add("Content-Type", "application/json")
	util2.LogError(err)

	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	responseString := string(responseBody)

	apiCallResult := util2.HttpError(response.StatusCode, responseString)
	if apiCallResult == false {
		os.Exit(0)
	} else if apiCallResult == true {
		jsonErr := json.Unmarshal(responseBody, &createUserResponse)
		util2.LogError(jsonErr)
		fmt.Println("User account successfully created for " + createUserResponse.Firstname + " " + createUserResponse.Lastname)
	}
}
