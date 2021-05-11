package update

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

func NewCmdAccessKeysUpdate() *cobra.Command {
	var (
		corsHeaders string
		disabled    bool
		id          string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates the properties of existing accessKey by accessId. It can be used to enable or disable the access key and to update the corsHeaders list.",
		Run: func(cmd *cobra.Command, args []string) {
			updateAccessKey(corsHeaders, disabled, id)
		},
	}
	cmd.Flags().StringVar(&corsHeaders, "corsHeaders", "", "Specify cors headers (they need to be comma separated e.g. header1,header2,header3")
	cmd.Flags().BoolVar(&disabled, "disabled", false, "Set to true to disable the access key")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the access key to update")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updateAccessKey(corsHeaders string, disabled bool, id string) {
	var accessKeyResponse api.GetAccessKeysResponse
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.UpdateAccessKeysRequest{
		Disabled: disabled,
	}
	if corsHeaders != "" {
		requestBodySchema.CorsHeaders = strings.Split(corsHeaders, ",")
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/accessKeys/" + id
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
