package list

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdSourceList() *cobra.Command {
	var (
		collectorId string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists sources assigned to a Sumo Logic collector",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			listSources(collectorId, log)
		},
	}

	cmd.Flags().StringVar(&collectorId, "collectorId", "", "Specify the collector id which the source is assigned to")
	cmd.MarkFlagRequired("collectorId")
	return cmd
}

func listSources(collectorId string, log zerolog.Logger) {
	var sourcesInfo api.ListSources

	requestUrl := "v1/collectors/" + collectorId + "/sources"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}

	jsonErr := json.Unmarshal(responseBody, &sourcesInfo)
	if jsonErr != nil {
		log.Error().Err(jsonErr).Msg("error unmarshalling response body")
	}

	sourcesInfoJson, err := json.MarshalIndent(sourcesInfo.Sources, "", "    ")

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Print(string(sourcesInfoJson))
	}
}
