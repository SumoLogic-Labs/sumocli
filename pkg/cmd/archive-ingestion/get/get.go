package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"net/url"
	"strconv"
)

func NewCmdArchiveIngestionGet() *cobra.Command {
	var (
		limit    int
		sourceId string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a list of all the ingestion jobs created on an Archive Source.",
		Run: func(cmd *cobra.Command, args []string) {
			getArchiveIngestion(limit, sourceId)
		},
	}
	cmd.Flags().IntVar(&limit, "limit", 10, "Specify the number of jobs to return")
	cmd.Flags().StringVar(&sourceId, "sourceId", "", "Specify the id of the Archive Source")
	cmd.MarkFlagRequired("sourceId")
	return cmd
}

func getArchiveIngestion(limit int, sourceId string) {
	var archiveIngestionResponse api.GetArchiveIngestion
	log := logging.GetConsoleLogger()
	requestUrl := "v1/archive/" + sourceId + "/jobs"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
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

	err = json.Unmarshal(responseBody, &archiveIngestionResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	archiveIngestionResponseJson, err := json.MarshalIndent(archiveIngestionResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(archiveIngestionResponseJson))
	}
}
