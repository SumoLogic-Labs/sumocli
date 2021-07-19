package export_status

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdExportStatus() *cobra.Command {
	var (
		contentId   string
		jobId       string
		isAdminMode bool
	)

	cmd := &cobra.Command{
		Use:   "export-status",
		Short: "Get the status of an asynchronous content export request for the given job identifier",
		Run: func(cmd *cobra.Command, args []string) {
			exportStatus(contentId, jobId, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&contentId, "contentId", "", "Specify the id of the content item to export")
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the export (returned from running sumocli content start-export)")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("contentId")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func exportStatus(contentId string, jobId string, isAdminMode bool) {
	var exportStatusResponse api.ExportStatusResponse
	log := logging.GetConsoleLogger()
	requestUrl := "/v2/content/" + contentId + "/export/" + jobId + "/status"
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

	err = json.Unmarshal(responseBody, &exportStatusResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	exportStatusJson, err := json.MarshalIndent(exportStatusResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal exportStatusResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(exportStatusJson))
	}
}
