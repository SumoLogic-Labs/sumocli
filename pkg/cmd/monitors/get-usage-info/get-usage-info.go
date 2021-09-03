package get_usage_info

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/api"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Incubator/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdMonitorsGetUsageInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-usage-info",
		Short: "Get the current number and the allowed number of log and metrics monitors.",
		Run: func(cmd *cobra.Command, args []string) {
			getUsageInfo()
		},
	}
	return cmd
}

func getUsageInfo() {
	var usageInfoResponse []api.GetUsageInfo
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/monitors/usageInfo"
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

	err = json.Unmarshal(responseBody, &usageInfoResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	usageInfoResponseJson, err := json.MarshalIndent(usageInfoResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(usageInfoResponseJson))
	}
}
