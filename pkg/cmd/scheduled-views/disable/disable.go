package disable

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdScheduledViewsDisable() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disable a scheduled view with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			disableScheduledView(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the scheduled view")
	cmd.MarkFlagRequired("id")
	return cmd
}

func disableScheduledView(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/scheduledViews/" + id + "/disable"
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		err := json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
	} else {
		fmt.Println("The scheduled view was disabled successfully.")
	}
}
