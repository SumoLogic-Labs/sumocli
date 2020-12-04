package change

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io/ioutil"
)

func NewCmdUserChangeEmail() *cobra.Command {
	var (
		id    string
		email string
	)

	cmd := &cobra.Command{
		Use:   "change email",
		Short: "Changes the email address of a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User change email request started.")
			userChangeEmail(id, email, logger)
			logger.Debug().Msg("User change email request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user that needs to have the email changed.")
	cmd.Flags().StringVar(&email, "email", "", "Specify the users new email address.")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("email")

	return cmd
}

func userChangeEmail(id string, email string, logger zerolog.Logger) {
	requestBodySchema := &api.UpdateUserEmail{
		Email: email,
	}
	requestBody, _ := json.Marshal(requestBodySchema)
	requestUrl := "v1/users/" + id + "/email/requestChange"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	logging.LogError(err, logger)

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		jsonErr := json.Unmarshal(responseBody, &responseError)
		logging.LogError(jsonErr, logger)
		if responseError.Errors[0].Code == "um1:unverified_email" {
			fmt.Println(responseError.Errors[0].Message)
		}
	} else {
		fmt.Println("Users email address successfully updated to: " + email)
	}
}
