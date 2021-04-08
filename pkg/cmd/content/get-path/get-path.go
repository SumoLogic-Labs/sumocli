package get_path

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdGetPath() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get-path",
		Short: "Gets the full path of a content item with the given identifier",
		Run: func(cmd *cobra.Command, args []string) {
			getPath(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getPath(id string) {
	var pathResponse api.GetPathResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/" + id + "/path"
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

	err = json.Unmarshal(responseBody, &pathResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	pathJson, err := json.MarshalIndent(pathResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal pathResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(pathJson))
	}
}
