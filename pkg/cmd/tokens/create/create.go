package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdTokensCreate() *cobra.Command {
	var (
		description string
		inactive    bool
		name        string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a token in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			createToken(description, inactive, name)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the token")
	cmd.Flags().BoolVar(&inactive, "inactive", false, "Set to true if you want the token to be inactive")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the token")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createToken(description string, inactive bool, name string) {
	var createTokenResponse api.GetTokenResponse
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateTokenRequest{
		Name:        name,
		Description: description,
		Type:        "CollectorRegistration",
	}
	if inactive == false {
		requestBodySchema.Status = "Active"
	} else {
		requestBodySchema.Status = "Inactive"
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/tokens"
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

	err = json.Unmarshal(responseBody, &createTokenResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	createTokenResponseJson, err := json.MarshalIndent(createTokenResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(createTokenResponseJson))
	}
}
