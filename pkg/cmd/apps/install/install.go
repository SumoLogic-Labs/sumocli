package install

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdAppsInstall() *cobra.Command {
	var (
		destinationFolderId string
		description         string
		logSource           string
		name                string
		uuid                string
	)

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs the app with given UUID in the folder specified.",
		Run: func(cmd *cobra.Command, args []string) {
			installApp(destinationFolderId, description, logSource, name, uuid)
		},
	}
	cmd.Flags().StringVar(&destinationFolderId, "destinationFolderId", "", "Specify the folder id that the app should be installed into")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the app")
	cmd.Flags().StringVar(&logSource, "logSource", "", "Specify a log source name (for example _sourceCategory=test)")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the app")
	cmd.Flags().StringVar(&uuid, "uuid", "", "Specify the UUID of the app to install")
	cmd.MarkFlagRequired("destinationFolderId")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("uuid")
	return cmd
}

func installApp(destinationFolderId string, description string, logSource string, name string, uuid string) {
	var installAppResponse api.InstallAppResponse
	log := logging.GetConsoleLogger()
	requestBodySchema := api.InstallAppRequest{
		Name:                name,
		Description:         description,
		DestinationFolderId: destinationFolderId,
		DataSourceValues: api.DataSourceValue{
			LogSrc: logSource,
		},
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/apps/" + uuid + "/install"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &installAppResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	installAppResponseJson, err := json.MarshalIndent(installAppResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(installAppResponseJson))
	}
}
