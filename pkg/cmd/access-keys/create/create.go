package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"strings"
)

func NewCmdAccessKeysCreate() *cobra.Command {
	var (
		name        string
		corsHeaders string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new access ID and key pair. The new access key can be used from the domains specified in corsHeaders field.",
		Run: func(cmd *cobra.Command, args []string) {
			createAccessKey(name, corsHeaders)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the access key")
	cmd.Flags().StringVar(&corsHeaders, "corsHeaders", "", "Specify cors headers (they need to be comma separated e.g. header1,header2,header3")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createAccessKey(name string, corsHeaders string) {
	var accessKeyResponse api.GetAccessKeysResponse
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateAccessKeysRequest{
		Label: name,
	}
	if corsHeaders != "" {
		requestBodySchema.CorsHeaders = strings.Split(corsHeaders, ",")
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/accessKeys"
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

	err = json.Unmarshal(responseBody, &accessKeyResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	accessKeyResponseJson, err := json.MarshalIndent(accessKeyResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(accessKeyResponseJson))
	}
}
