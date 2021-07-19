package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdDashboardsGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a dashboard by the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getDashboards(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the dashboard")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getDashboards(id string) {
	var dashboardResponse api.GetDashboard
	log := logging.GetConsoleLogger()
	requestUrl := "/v2/dashboards/" + id
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
