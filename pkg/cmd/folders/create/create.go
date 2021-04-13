package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdCreate() *cobra.Command {
	var (
		name        string
		description string
		parentId    string
		isAdminMode bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new folder under the given parent folder.",
		Run: func(cmd *cobra.Command, args []string) {
			create(name, description, parentId, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the folder")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the folder")
	cmd.Flags().StringVar(&parentId, "parentId", "", "Specify the parent folder Id")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("parentId")
	return cmd
}

func create(name string, description string, parentId string, isAdminMode bool) {
	var foldersResponse api.FolderResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders"
	requestBodySchema := &api.CreateFolderRequest{
		Name:        name,
		Description: description,
		ParentId:    parentId,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal requestBody")
	}
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
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

	createFoldersResponseJson, err := json.MarshalIndent(foldersResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal foldersResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(createFoldersResponseJson))
	}
}
