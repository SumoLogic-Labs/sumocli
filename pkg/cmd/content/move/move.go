package move

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

func NewCmdMove() *cobra.Command {
	var (
		id                  string
		destinationFolderId string
		isAdminMode         bool
	)

	cmd := &cobra.Command{
		Use:   "move",
		Short: "Moves an item from its current location to another folder.",
		Run: func(cmd *cobra.Command, args []string) {
			move(id, destinationFolderId, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content to move")
	cmd.Flags().StringVar(&destinationFolderId, "destinationFolderId", "", "Specify the destination folder to move the content to")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("destinationFolderId")
	return cmd
}

func move(id string, destinationFolderId string, isAdminMode bool) {
	var moveResponse api.MoveResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/" + id + "/move"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	if isAdminMode == true {
		request.Header.Add("isAdminMode", "true")
	}
	query := url.Values{}
	query.Add("destinationFolderId", destinationFolderId)
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

	err = json.Unmarshal(responseBody, &moveResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	moveJson, err := json.MarshalIndent(moveResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal copyResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
		fmt.Println(string(moveJson))
	} else {
		fmt.Println("Content successfully moved.")
	}
}
