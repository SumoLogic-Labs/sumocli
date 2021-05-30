package get_capacity_info

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldManagementGetCapacityInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-capacity-info",
		Short: "Every account has a limited number of fields available." +
			"This command returns your account limitations and remaining quota",
		Run: func(cmd *cobra.Command, args []string) {
			getCapacityInfo()
		},
	}
	return cmd
}

func getCapacityInfo() {
	var capacityResponse api.GetCapacityInformation
	log := logging.GetConsoleLogger()
	requestUrl := "v1/fields/quota"
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

	err = json.Unmarshal(responseBody, &capacityResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	capacityResponseJson, err := json.MarshalIndent(capacityResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(capacityResponseJson))
	}
}
