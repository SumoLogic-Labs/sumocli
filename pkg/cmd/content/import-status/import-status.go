package import_status

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdImportStatus() *cobra.Command {
	var (
		folderId    string
		jobId       string
		isAdminMode bool
	)

	cmd := &cobra.Command{
		Use:   "import-status",
		Short: "Get the status of an asynchronous content import request for the given job identifier",
		Run: func(cmd *cobra.Command, args []string) {
			importStatus(folderId, jobId, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&folderId, "folderId", "", "Specify the id of the folder to import to")
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the import (returned from running sumocli content start-import)")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("folderId")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func importStatus(folderId string, jobId string, isAdminMode bool) {
	var importStatusResponse api.ExportStatusResponse
	log := logging.GetConsoleLogger()
	requestUrl := "/v2/content/folders/" + folderId + "/import/" + jobId + "/status"
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

	err = json.Unmarshal(responseBody, &importStatusResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	importStatusJson, err := json.MarshalIndent(importStatusResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal exportStatusResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(importStatusJson))
	}
}
