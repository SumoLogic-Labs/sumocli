package enable_lockdown

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/api"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Labs/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdSamlEnableLockdown() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable-lockdown",
		Short: "Enabling SAML lockdown requires users to sign in using SAML preventing them from logging in with an email and password.",
		Run: func(cmd *cobra.Command, args []string) {
			enableSamlLockdown()
		},
	}
	return cmd
}

func enableSamlLockdown() {
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/saml/lockdown/enable"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		err := json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
	} else {
		fmt.Println("SAML lockdown was enabled successfully.")
	}
}
