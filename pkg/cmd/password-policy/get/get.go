package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdPasswordPolicyGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the current password policy.",
		Run: func(cmd *cobra.Command, args []string) {
			getPasswordPolicy()
		},
	}
	return cmd
}

func getPasswordPolicy() {
	var passwordPolicyResponse api.GetPasswordPolicy
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/passwordPolicy"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &passwordPolicyResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	passwordPolicyResponseJson, err := json.MarshalIndent(passwordPolicyResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(passwordPolicyResponseJson))
	}
}
