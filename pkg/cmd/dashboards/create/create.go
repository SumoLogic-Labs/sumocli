package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"os"
)

func NewCmdDashboardsCreate() *cobra.Command {
	var file string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new dashboard.",
		Run: func(cmd *cobra.Command, args []string) {
			createDashboard(file)
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "Specify the full file path to a json file containing a dashboard definition."+
		"The definition can be retrieved from running sumocli dashboards get or from exporting the dashboard in the UI.")
	cmd.MarkFlagRequired("file")
	return cmd
}

func createDashboard(file string) {
	var dashboardResponse api.GetDashboard
	var dashboardRequest api.CreateDashboard
	log := logging.GetConsoleLogger()
	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Error().Err(err).Msg("failed to read file " + file)
	}
	err = json.Unmarshal(fileData, &dashboardRequest)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal file")
	}
	requestBody, err := json.Marshal(dashboardRequest)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "/v2/dashboards"
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

	err = json.Unmarshal(responseBody, &dashboardResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	dashboardResponseJson, err := json.MarshalIndent(dashboardResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal foldersResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(dashboardResponseJson))
	}
}
