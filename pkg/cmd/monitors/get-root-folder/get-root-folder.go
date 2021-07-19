package get_root_folder

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdMonitorsGetRootFolder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-root-folder",
		Short: "Get the root folder in the monitors library.",
		Run: func(cmd *cobra.Command, args []string) {
			getRootFolder()
		},
	}
	return cmd
}

func getRootFolder() {
	var rootFolderRepsonse api.RootMonitorFolder
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/monitors/root"
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

	err = json.Unmarshal(responseBody, &rootFolderRepsonse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	rootFolderResponseJson, err := json.MarshalIndent(rootFolderRepsonse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(rootFolderResponseJson))
	}
}
