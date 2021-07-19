package list

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdServiceAllowlistList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all allowlisted CIDR notations and/or IP addresses for the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			listServiceAllowlist()
		},
	}
	return cmd
}

func listServiceAllowlist() {
	var allowlistResponse api.ListServiceAllowlist
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/serviceAllowlist/addresses"
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

	err = json.Unmarshal(responseBody, &allowlistResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	allowlistResponseJson, err := json.MarshalIndent(allowlistResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(allowlistResponseJson))
	}
}
