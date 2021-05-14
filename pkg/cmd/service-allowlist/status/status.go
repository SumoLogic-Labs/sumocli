package status

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdServiceAllowlistStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status of the service allowlisting functionality for login/API authentication or content sharing for the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getServiceAllowlistStatus()
		},
	}
	return cmd
}

func getServiceAllowlistStatus() {
	var allowlistStatusResponse api.GetAllowlistStatus
	log := logging.GetConsoleLogger()
	requestUrl := "v1/serviceAllowlist/status"
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

	err = json.Unmarshal(responseBody, &allowlistStatusResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	allowlistStatusResponseJson, err := json.MarshalIndent(allowlistStatusResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(allowlistStatusResponseJson))
	}
}
