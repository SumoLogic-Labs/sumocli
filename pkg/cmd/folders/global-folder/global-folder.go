package global_folder

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdGlobalFolder() *cobra.Command {
	var isAdminMode bool

	cmd := &cobra.Command{
		Use: "global-folder",
		Short: "Schedule an asynchronous job to get global folder. " +
			"Global folder contains all content items that a user has permissions to view in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			globalFolder(isAdminMode)
		},
	}
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	return cmd
}

func globalFolder(isAdminMode bool) {
	var globalFolderResponse api.GlobalFolderResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders/global"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	if isAdminMode == true {
		request.Header.Add("isAdminMode", "true")
	}
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &globalFolderResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	globalFoldersResponseJson, err := json.MarshalIndent(globalFolderResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal foldersResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(globalFoldersResponseJson))
	}
}
