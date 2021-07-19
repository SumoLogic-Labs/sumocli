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

func NewCmdPartitionsGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a partition with the given identifier from the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			getPartition(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the partition")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getPartition(id string) {
	var partitionsResponse api.Partitions
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/partitions/" + id
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

	err = json.Unmarshal(responseBody, &partitionsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	partitionsResponseJson, err := json.MarshalIndent(partitionsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(partitionsResponseJson))
	}
}
