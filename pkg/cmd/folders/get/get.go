package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdGet() *cobra.Command {
	var (
		id          string
		isAdminMode bool
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a folder with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			get(id, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the identifier of the folder")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	return cmd
}

func get(id string, isAdminMode bool) {
	var foldersResponse api.FolderResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders/" + id
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

	err = json.Unmarshal(responseBody, &foldersResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	foldersResponseJson, err := json.MarshalIndent(foldersResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal foldersResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(foldersResponseJson))
	}
}
