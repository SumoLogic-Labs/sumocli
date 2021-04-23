package admin_recommended_folder

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdAdminRecommendedFolder() *cobra.Command {
	var isAdminMode bool

	cmd := &cobra.Command{
		Use:   "admin-recommended-folder",
		Short: "Schedule an asynchronous job to get the top-level Admin Recommended content items.",
		Run: func(cmd *cobra.Command, args []string) {
			adminRecommendedFolder(isAdminMode)
		},
	}
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	return cmd
}

func adminRecommendedFolder(isAdminMode bool) {
	var adminRecommendedFolderResponse api.GlobalFolderResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders/adminRecommended"
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

	err = json.Unmarshal(responseBody, &adminRecommendedFolderResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	adminRecommendedFolderResponseJson, err := json.MarshalIndent(adminRecommendedFolderResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal foldersResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(adminRecommendedFolderResponseJson))
	}
}
