package start_export

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdStartExport() *cobra.Command {
	var (
		id          string
		isAdminMode bool
	)

	cmd := &cobra.Command{
		Use: "start-export",
		Short: "Starts an asynchronous export of content with the given identifier. You will be given a job identifier" +
			"which can be used with the sumocli content export-status command." +
			"If the content is a folder everything under that folder is exported recursively.",
		Run: func(cmd *cobra.Command, args []string) {
			startExport(id, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content item to export")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	return cmd
}

func startExport(id string, isAdminMode bool) {
	var exportResponse api.StartExportResponse
	log := logging.GetConsoleLogger()
	requestUrl := "/v2/content/" + id + "/export"
	client, request := factory.NewHttpRequest("POST", requestUrl)
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

	err = json.Unmarshal(responseBody, &exportResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	exportJson, err := json.MarshalIndent(exportResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal exportResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(exportJson))
	}
}
