package get_configurations

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/api"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Incubator/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdSamlGetConfigurations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-configurations",
		Short: "Get a list of all SAML configurations in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getSaml()
		},
	}
	return cmd
}

func getSaml() {
	var samlResponse []api.GetSaml
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/saml/identityProviders"
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

	err = json.Unmarshal(responseBody, &samlResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	samlResponseJson, err := json.MarshalIndent(samlResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(samlResponseJson))
	}
}
