package builds

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdGetBuilds() *cobra.Command {
	var ()

	cmd := &cobra.Command{
		Use:   "builds",
		Short: "Gets available Sumo Logic collector builds",
		Run: func(cmd *cobra.Command, args []string) {
			getBuilds()
		},
	}

	return cmd
}

func getBuilds() {
	log := logging.GetConsoleLogger()
	var builds api.Targets
	requestUrl := "v1/collectors/upgrades/targets"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}

	err = json.Unmarshal(responseBody, &builds)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}

	buildsJson, err := json.MarshalIndent(builds, "", "    ")

	if response.StatusCode == 200 {
		fmt.Println(string(buildsJson))
	} else {
		factory.HttpError(response.StatusCode, responseBody, log)
	}
}
