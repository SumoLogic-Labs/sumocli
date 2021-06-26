package disable

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"os"
)

func NewCmdUserDisableMFA() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "disable mfa",
		Short: "Disables MFA for a Sumo Logic user (this command only works interactively).",
		Run: func(cmd *cobra.Command, args []string) {
			userDisableMFA()
		},
	}

	return cmd
}

func userDisableMFA() {
	log := logging.GetConsoleLogger()
	validate := func(input string) error {
		if input == "" {
			return errors.New("Value is empty")
		}
		return nil
	}

	promptId := promptui.Prompt{
		Label:    "Please enter the Sumo Logic id for the user",
		Validate: validate,
	}

	promptEmail := promptui.Prompt{
		Label:    "Please enter the Sumo Logic users email address",
		Validate: validate,
	}

	promptPassword := promptui.Prompt{
		Label:    "Please enter the Sumo Logic users password",
		Mask:     '*',
		Validate: validate,
	}

	promptConfirm := promptui.Prompt{
		Label:     "Confirm that you want to disable MFA? Removing MFA can be a security risk!",
		IsConfirm: true,
	}

	idResult, err := promptId.Run()
	emailResult, err := promptEmail.Run()
	passwordResult, err := promptPassword.Run()
	_, err = promptConfirm.Run()

	if err != nil {
		log.Error().Err(err).Msg("failed to generate prompt")
		os.Exit(0)
	}

	requestBodySchema := &api.DisableUserMfa{
		Email:    emailResult,
		Password: passwordResult,
	}
	requestBody, _ := json.Marshal(requestBodySchema)
	requestUrl := "v1/users/" + idResult + "/mfa/disable"
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

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		err = json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		if responseError.Errors[0].Message != "" {
			fmt.Println(responseError.Errors[0].Message)
		} else if responseError.Errors[0].Code == "auth1:mfa_not_allowed" {
			fmt.Println("MFA is not enabled on user " + emailResult)
		}
	} else {
		fmt.Println("MFA removed from user " + emailResult)
	}
}
