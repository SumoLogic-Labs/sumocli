package delete_configuration

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/api"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Incubator/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdSamlDeleteConfiguration() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete-configuration",
		Short: "Delete a SAML configuration with the given identifier from the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteSamlConfiguration(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the SAML configuration")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteSamlConfiguration(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/saml/identityProviders/" + id
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
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
		fmt.Println("The SAML configuration was deleted successfully.")
	}
}
