package list_personal

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdAccessKeysListPersonal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-personal",
		Short: "List all access keys that belong to your user.",
		Run: func(cmd *cobra.Command, args []string) {
			listPersonalAccessKeys()
		},
	}
	return cmd
}

func listPersonalAccessKeys() {
	var accessKeyResponse api.ListAccessKeysResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v1/accessKeys/personal"
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
