package admin_recommended_folder_status

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdAdminRecommendedFolderStatus() *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "admin-recommended-folder-status",
		Short: "Get the status of an asynchronous Admin Recommended folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			adminRecommendedFolderStatus(jobId)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id (returned from running sumocli admin-recommended-folder)")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func adminRecommendedFolderStatus(jobId string) {
	var adminRecommendedFolderStatusResponse api.GlobalFolderStatusResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders/adminRecommended/" + jobId + "/status"
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

	err = json.Unmarshal(responseBody, &adminRecommendedFolderStatusResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	adminRecommendedFolderStatusResponseJson, err := json.MarshalIndent(adminRecommendedFolderStatusResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal exportStatusResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(adminRecommendedFolderStatusResponseJson))
	}
}
