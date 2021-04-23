package global_folder_result

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdGlobalFolderResult() *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "global-folder-result",
		Short: "Get results from global folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			globalFolderResult(jobId)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the global folder (returned from running sumocli folders global-folder")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func globalFolderResult(jobId string) {
	var globalFolderResultResponse api.GlobalFolderResultResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders/global/" + jobId + "/result"
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

	err = json.Unmarshal(responseBody, &globalFolderResultResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	globalFolderResultResponseJson, err := json.MarshalIndent(globalFolderResultResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal exportStatusResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(globalFolderResultResponseJson))
	}
}
