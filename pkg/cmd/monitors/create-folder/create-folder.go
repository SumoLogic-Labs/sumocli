package create_folder

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

func NewCmdMonitorsCreateFolder() *cobra.Command {
	var (
		description string
		name        string
		parentId    string
	)

	cmd := &cobra.Command{
		Use:   "create-folder",
		Short: "Create a folder in the monitors library",
		Run: func(cmd *cobra.Command, args []string) {
			createFolder(description, name, parentId)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify the description for the monitor folder")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name for the monitor folder")
	cmd.Flags().StringVar(&parentId, "parentId", "", "Specify the parent Id of the monitor folder. "+
		"You can get the root monitor folder id by running sumocli monitors get-root-folder")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createFolder(description string, name string, parentId string) {
	var monitorFolderResponse api.MonitorFolder
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateMonitorFolder{
		Description: description,
		Name:        name,
		Type:        "MonitorsLibraryFolder",
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/monitors"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	query := url.Values{}
	query.Add("parentId", parentId)
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &monitorFolderResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	monitorFolderResponseJson, err := json.MarshalIndent(monitorFolderResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(monitorFolderResponseJson))
	}
}
