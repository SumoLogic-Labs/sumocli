package disable_lockdown

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdSamlDisableLockdown() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable-lockdown",
		Short: "Disable SAML lockdown for the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			disableSamlLockdown()
		},
	}
	return cmd
}

func disableSamlLockdown() {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/saml/lockdown/disable"
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
		fmt.Println("SAML lockdown was disabled successfully.")
	}
}
