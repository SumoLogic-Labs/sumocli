package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdUpdate() *cobra.Command {
	var (
		name        string
		description string
		id          string
		isAdminMode bool
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing folder with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			update(name, description, id, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the folder")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the folder")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the folder to update")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("id")
	return cmd
}

func update(name string, description string, id string, isAdminMode bool) {
	var foldersResponse api.FolderResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders/" + id
	requestBodySchema := &api.UpdateFolderRequest{
		Name:        name,
		Description: description,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal requestBody")
	}
	client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
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

	updateFoldersResponseJson, err := json.MarshalIndent(foldersResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal foldersResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(updateFoldersResponseJson))
	}
}
