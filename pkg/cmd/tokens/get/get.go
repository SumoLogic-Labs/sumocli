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

func NewCmdTokensGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a token with the given identifier in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			getToken(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the token to retrieve")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getToken(id string) {
	var tokenResponse api.GetTokenResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v1/tokens/" + id
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

	err = json.Unmarshal(responseBody, &tokenResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	tokenResponseJson, err := json.MarshalIndent(tokenResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(tokenResponseJson))
	}
}
