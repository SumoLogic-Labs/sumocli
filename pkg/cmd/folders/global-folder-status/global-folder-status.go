package global_folder_status

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdGlobalFolderStatus() *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "global-folder-status",
		Short: "Get the status of an asynchronous global folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			globalFolderStatus(jobId)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the global folder (returned from running sumocli folders global-folder")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func globalFolderStatus(jobId string) {
	var globalFolderStatusResponse api.GlobalFolderStatusResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders/global/" + jobId + "/status"
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

	err = json.Unmarshal(responseBody, &globalFolderStatusResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	globalFolderStatusResponseJson, err := json.MarshalIndent(globalFolderStatusResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal exportStatusResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(globalFolderStatusResponseJson))
	}
}
