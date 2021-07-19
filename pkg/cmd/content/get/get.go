package get

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

func NewCmdGet() *cobra.Command {
	var path string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a content item corresponding to the provided path",
		Run: func(cmd *cobra.Command, args []string) {
			getContent(path)
		},
	}
	cmd.Flags().StringVar(&path, "path", "", "Specify the path of the content you want to retrieve (e.g. /Library/Users/user@demo.com/SampleFolder)")
	cmd.MarkFlagRequired("path")
	return cmd
}

func getContent(path string) {
	var contentResponse api.GetContentResponse
	log := logging.GetConsoleLogger()
	requestUrl := "/v2/content/path"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	query.Add("path", path)
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &contentResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	contentJson, err := json.MarshalIndent(contentResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal contentResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(contentJson))
	}
}
