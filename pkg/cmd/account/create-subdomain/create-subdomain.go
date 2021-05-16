package create_subdomain

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdAccountCreateSubdomain() *cobra.Command {
	var subdomain string

	cmd := &cobra.Command{
		Use:   "create-subdomain",
		Short: "Create a subdomain. Only the Account Owner can create a subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			createSubdomain(subdomain)
		},
	}
	cmd.Flags().StringVar(&subdomain, "subdomain", "", "Specify a subdomain (minimum 4 and maximum 63 characters)")
	cmd.MarkFlagRequired("subdomain")
	return cmd
}

func createSubdomain(subdomain string) {
	var subdomainResponse api.GetSubdomain
	log := logging.GetConsoleLogger()
	requestBodySchema := api.UpdateSubdomainRequest{
		Subdomain: subdomain,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/account/subdomain"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &subdomainResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	subdomainResponseJson, err := json.MarshalIndent(subdomainResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(subdomainResponseJson))
	}
}
