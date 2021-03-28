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

func NewCmdCollectorList() *cobra.Command {
	var (
		filter  string
		limit   int
		offset  string
		offline bool
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic collectors",
		Run: func(cmd *cobra.Command, args []string) {
			listCollectors(filter, limit, offset, offline)
		},
	}

	cmd.Flags().StringVar(&filter, "filter", "", "Filters the collectors returned using either installed, hosted, dead or alive")
	cmd.Flags().IntVar(&limit, "limit", 1000, "Maximum number of collectors returned")
	cmd.Flags().StringVar(&offset, "offset", "", "Offset into the list of collectors")
	cmd.Flags().BoolVar(&offline, "offline", false, "Lists offline collectors")
	return cmd
}

func listCollectors(filter string, limit int, offset string, offline bool) {
	log := logging.GetConsoleLogger()
	var collectorInfo api.Collectors
	var requestUrl string
	if offline == true {
		requestUrl = "v1/collectors/offline"
	} else {
		requestUrl = "v1/collectors"
	}

	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	if filter != "" && offline == false {
		if factory.ValidateCollectorFilter(filter) == false {
			log.Fatal().Msg(filter + "is an invalid field to filter by. Available fields are installed, hosted, dead or alive.")
		} else {
			query.Add("filter", filter)
		}
	}
	if offset != "" && offline == false {
		query.Add("offset", offset)
	}
	query.Add("limit", strconv.Itoa(limit))
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}

	jsonErr := json.Unmarshal(responseBody, &collectorInfo)
	if jsonErr != nil {
		log.Error().Err(jsonErr).Msg("error unmarshalling response body")
	}

	collectorInfoJson, err := json.MarshalIndent(collectorInfo.Data, "", "    ")

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Print(string(collectorInfoJson))
	}
}
