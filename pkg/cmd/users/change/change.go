package change

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
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
			userChangeEmail(id, email)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user that needs to have the email changed.")
	cmd.Flags().StringVar(&email, "email", "", "Specify the users new email address.")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("email")

	return cmd
}

func userChangeEmail(id string, email string) {
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.UpdateUserEmail{
		Email: email,
	}
	requestBody, _ := json.Marshal(requestBodySchema)
	requestUrl := "v1/users/" + id + "/email/requestChange"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		err = json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		if responseError.Errors[0].Code == "um1:unverified_email" {
			fmt.Println(responseError.Errors[0].Message)
		}
	} else {
		fmt.Println("Users email address successfully updated to: " + email)
	}
}
