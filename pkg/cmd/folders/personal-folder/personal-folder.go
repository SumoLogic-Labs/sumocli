package personal_folder

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdPersonalFolder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "personal-folder",
		Short: "Get the personal folder of the current user.",
		Run: func(cmd *cobra.Command, args []string) {
			personalFolder()
		},
	}
	return cmd
}

func personalFolder() {
	var foldersResponse api.FolderResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders/personal"
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

	err = json.Unmarshal(responseBody, &foldersResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	personalFoldersResponseJson, err := json.MarshalIndent(foldersResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal foldersResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(personalFoldersResponseJson))
	}
}
