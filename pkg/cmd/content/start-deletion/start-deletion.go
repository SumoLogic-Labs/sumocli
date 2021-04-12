package start_deletion

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdStartDeletion() *cobra.Command {
	var (
		id          string
		isAdminMode bool
	)

	cmd := &cobra.Command{
		Use:   "start-deletion",
		Short: "Start an asynchronous content deletion job with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			startDeletion(id, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content to delete")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	return cmd
}

func startDeletion(id string, isAdminMode bool) {
	var deletionResponse api.StartExportResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/" + id + "/delete"
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
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

	err = json.Unmarshal(responseBody, &deletionResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	exportJson, err := json.MarshalIndent(deletionResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal exportResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(exportJson))
	}
}
