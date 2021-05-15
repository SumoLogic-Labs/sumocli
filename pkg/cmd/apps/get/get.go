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

func NewCmdAppsGet() *cobra.Command {
	var uuid string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets the app with the given universally unique identifier (UUID).",
		Run: func(cmd *cobra.Command, args []string) {
			getApp(uuid)
		},
	}
	cmd.Flags().StringVar(&uuid, "uuid", "", "Specify the UUID of the app")
	cmd.MarkFlagRequired("uuid")
	return cmd
}

func getApp(uuid string) {
	var appResponse api.AppDetails
	log := logging.GetConsoleLogger()
	requestUrl := "v1/apps/" + uuid
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

	err = json.Unmarshal(responseBody, &appResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	appResponseJson, err := json.MarshalIndent(appResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(appResponseJson))
	}
}
