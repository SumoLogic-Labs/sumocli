package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdTokensUpdate() *cobra.Command {
	var (
		description string
		id          string
		inactive    bool
		name        string
		version     int
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a token with the given identifier in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			updateToken(description, id, inactive, name, version)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the token")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the token to update")
	cmd.Flags().BoolVar(&inactive, "inactive", false, "Set to true if you want the token to be inactive")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the token")
	cmd.Flags().IntVar(&version, "version", 0, "Specify a version of the token")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("version")
	return cmd
}

func updateToken(description string, id string, inactive bool, name string, version int) {
	var updateTokenResponse api.GetTokenResponse
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.UpdateTokenRequest{
		Name:        name,
		Description: description,
		Type:        "CollectorRegistration",
		Version:     0,
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
	requestUrl := "v1/tokens/" + id
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

	err = json.Unmarshal(responseBody, &updateTokenResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	updateTokenResponseJson, err := json.MarshalIndent(updateTokenResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(updateTokenResponseJson))
	}
}
