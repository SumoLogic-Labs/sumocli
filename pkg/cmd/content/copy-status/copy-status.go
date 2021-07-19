package copy_status

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdCopyStatus() *cobra.Command {
	var (
		id          string
		jobId       string
		isAdminMode bool
	)

	cmd := &cobra.Command{
		Use:   "copy-status",
		Short: "Get the status of the copy request with the given job identifier. On success, field statusMessage will contain identifier of the newly copied content.",
		Run: func(cmd *cobra.Command, args []string) {
			copyStatus(id, jobId, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content that was copied")
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the copy (returned from running sumocli content start-copy)")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func copyStatus(id string, jobId string, isAdminMode bool) {
	var copyStatusResponse api.ExportStatusResponse
	log := logging.GetConsoleLogger()
	requestUrl := "/v2/content/" + id + "/copy/" + jobId + "/status"
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

	err = json.Unmarshal(responseBody, &copyStatusResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	copyStatusJson, err := json.MarshalIndent(copyStatusResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal exportStatusResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(copyStatusJson))
	}
}
