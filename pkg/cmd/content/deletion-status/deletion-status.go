package deletion_status

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdDeletionStatus() *cobra.Command {
	var (
		id          string
		jobId       string
		isAdminMode bool
	)

	cmd := &cobra.Command{
		Use:   "deletion-status",
		Short: "Get the status of an asynchronous content deletion job request for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			deletionStatus(id, jobId, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content to delete")
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the deletion (returned from running sumocli content start-deletion)")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("contentId")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func deletionStatus(id string, jobId string, isAdminMode bool) {
	var deletionStatusResponse api.ExportStatusResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/" + id + "/delete/" + jobId + "/status"
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

	err = json.Unmarshal(responseBody, &deletionStatusResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	importStatusJson, err := json.MarshalIndent(deletionStatusResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal exportStatusResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(importStatusJson))
	}
}
