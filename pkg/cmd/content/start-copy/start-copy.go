package start_copy

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"net/url"
)

func NewCmdStartCopy() *cobra.Command {
	var (
		id                string
		destinationFolder string
		isAdminMode       bool
	)

	cmd := &cobra.Command{
		Use:   "start-copy",
		Short: "Start an asynchronous content copy job with the given identifier to the destination folder. If the content item is a folder, everything under the folder is copied recursively.",
		Run: func(cmd *cobra.Command, args []string) {
			startCopy(id, destinationFolder, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content you want to copy")
	cmd.Flags().StringVar(&destinationFolder, "destinationFolder", "", "Specify the id of the destination folder")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "et to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("destinationFolder")
	return cmd
}

func startCopy(id string, destinationFolder string, isAdminMode bool) {
	var copyResponse api.StartExportResponse
	log := logging.GetConsoleLogger()
	requestUrl := "/v2/content/" + id + "/copy"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	if isAdminMode == true {
		request.Header.Add("isAdminMode", "true")
	}
	query := url.Values{}
	query.Add("destinationFolder", destinationFolder)
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to ")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &copyResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	copyJson, err := json.MarshalIndent(copyResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal copyResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(copyJson))
	}
}
