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

func NewCmdGetUpgradableCollectors() *cobra.Command {
	var (
		toVersion string
		offset    int
		limit     int
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets collectors in Sumo Logic that are upgradable",
		Run: func(cmd *cobra.Command, args []string) {
			getUpgradableCollectors(toVersion, offset, limit)
		},
	}

	cmd.Flags().StringVar(&toVersion, "toVersion", "", "Collector build to upgrade to, if not specified defaults to the latest version")
	cmd.Flags().IntVar(&offset, "offset", 0, "Offset into the list of collectors")
	cmd.Flags().IntVar(&limit, "limit", 50, "Maximum number of collectors to return")
	return cmd
}

func getUpgradableCollectors(toVersion string, offset int, limit int) {
	log := logging.GetConsoleLogger()
	var collectorInfo api.Collectors
	requestUrl := "v1/collectors/upgrades/collectors"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	query.Add("toVersion", toVersion)
	query.Add("offset", strconv.Itoa(offset))
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

	err = json.Unmarshal(responseBody, &collectorInfo)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}

	collectorInfoJson, err := json.MarshalIndent(collectorInfo.Data, "", "    ")

	if response.StatusCode == 200 {
		fmt.Println(string(collectorInfoJson))
	} else {
		factory.HttpError(response.StatusCode, responseBody, log)
	}
}
