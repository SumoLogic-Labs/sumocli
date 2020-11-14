package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"github.com/wizedkyle/sumocli/pkg/logging"
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
		Long:  "Creates a Sumo Logic user by specifying the first name, last name, email and roleIds",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User create request started.")
			user(firstName, lastName, emailAddress, roleIds, logger)
			logger.Debug().Msg("User create request finished.")
		},
	}

	cmd.Flags().StringVar(&firstName, "firstname", "", "First name of the user")
	cmd.Flags().StringVar(&lastName, "lastname", "", "Last name of the user")
	cmd.Flags().StringVar(&emailAddress, "email", "", "Email address of the user")
	cmd.Flags().StringArrayVar(&roleIds, "roleids", []string{}, "Comma deliminated list of Role Ids")

	return cmd
}

func user(firstName string, lastName string, emailAddress string, roleIds []string, logger zerolog.Logger) {
	var createUserResponse CreateUserResponse
	client := util2.GetHttpClient()
	authToken, apiEndpoint := login.ReadCredentials()

	requestBodySchema := &CreateUserRequest{
		Firstname:    firstName,
		Lastname:     lastName,
		Emailaddress: emailAddress,
		Roleids:      roleIds,
	}

	requestBody, _ := json.Marshal(requestBodySchema)

	request, err := http.NewRequest("POST", apiEndpoint+"v1/users", bytes.NewBuffer(requestBody))
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	logging.LogErrorWithMessage("Creating authorization header failed, please review the credentials supplied in sumocli login.", err, logger)

	response, err := client.Do(request)
	logging.LogErrorWithMessage("Authorization was not successful, please review your connectivity and credentials.", err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	responseString := string(responseBody)

	apiCallResult := logging.HttpError(response.StatusCode, responseString)
	if apiCallResult == false {
		os.Exit(0)
	} else if apiCallResult == true {
		jsonErr := json.Unmarshal(responseBody, &createUserResponse)
		logging.LogError(jsonErr, logger)
		fmt.Println("User account successfully created for " + createUserResponse.Firstname + " " + createUserResponse.Lastname)
	}
}
