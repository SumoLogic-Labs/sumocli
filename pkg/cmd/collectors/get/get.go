package get

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

func NewCmdCollectorGet() *cobra.Command {
	var (
		id   string
		name string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic collector information",
		Long:  "You can use either the id or the name of the collector to specify the collector to return",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			getCollector(id, name, log)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the collector")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the collector")
	return cmd
}

func getCollector(id string, name string, log zerolog.Logger) {
	var collectorInfo api.CollectorResponse
	requestUrl := "v1/collectors/"
	if id != "" {
		requestUrl = requestUrl + id
	} else if name != "" {
		requestUrl = requestUrl + "name/" + name
	} else {
		log.Fatal().Msg("please specify either a id or name of a collector")
	}

	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("error reading response body from request")
	}

	jsonErr := json.Unmarshal(responseBody, &collectorInfo)
	if jsonErr != nil {
		log.Fatal().Err(err).Msg("error unmarshalling response body")
	}

	collectorInfoJson, err := json.MarshalIndent(collectorInfo, "", "    ")

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(collectorInfoJson))
	}
}
