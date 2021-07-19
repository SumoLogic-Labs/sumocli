package job_status

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdLookupTableJobStatus() *cobra.Command {
	var (
		jobId      string
		jsonFormat bool
	)

	cmd := &cobra.Command{
		Use:   "job-status",
		Short: "Retrieve the status of a previously made request using sumocli lookup-tables upload or sumocli lookup-tables empty",
		Run: func(cmd *cobra.Command, args []string) {
			getLookupTableJobStatus(jobId, jsonFormat)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the Job ID to get the status for (returned from running sumocli lookup-tables upload or sumocli lookup-tables empty")
	cmd.Flags().BoolVar(&jsonFormat, "jsonFormat", false, "Set to true if you want the output to be formatted JSON")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getLookupTableJobStatus(jobId string, jsonFormat bool) {
	var statusResponse api.LookupTableStatusResponse
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/lookupTables/jobs/" + jobId + "/status"
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

	err = json.Unmarshal(responseBody, &statusResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		if jsonFormat == true {
			statusResponseJson, err := json.MarshalIndent(statusResponse, "", "    ")
			if err != nil {
				log.Error().Err(err).Msg("failed to marshal response")
			}
			fmt.Print(string(statusResponseJson))
		} else {
			statusResponseJson, err := json.Marshal(statusResponse)
			if err != nil {
				log.Error().Err(err).Msg("failed to marshal response")
			}
			fmt.Println(string(statusResponseJson))
		}
	}
}
