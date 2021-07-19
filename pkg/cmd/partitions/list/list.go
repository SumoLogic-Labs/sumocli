package list

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

func NewCmdPartitionsList() *cobra.Command {
	var (
		limit     int
		viewTypes string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all partitions in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			listPartitions(limit, viewTypes)
		},
	}
	cmd.Flags().IntVar(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	cmd.Flags().StringVar(&viewTypes, "viewTypes", "", "Specify the type of partitions to retrieve. "+
		"Valid values are: DefaultView, Partition, and AuditIndex. You can add multiple types for example DefaultView,Partition.")
	return cmd
}

func listPartitions(limit int, viewTypes string) {
	var partitionsResponse api.GetPartitions
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/partitions"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	query.Add("viewTypes", viewTypes)
	request.URL.RawQuery = query.Encode()
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
