package list

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io/ioutil"
	"net/url"
	"strings"
)

func NewCmdCollectorList() *cobra.Command {
	var (
		filter  string
		limit   string
		offset  string
		output  string
		offline bool
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic collectors",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			listCollectors(filter, limit, offset, output, offline, log)
		},
	}

	cmd.Flags().StringVar(&filter, "filter", "", "Filters the collectors returned using either installed, hosted, dead or alive")
	cmd.Flags().StringVar(&limit, "limit", "", "Maximum number of collectors returned")
	cmd.Flags().StringVar(&offset, "offset", "", "Offset into the list of collectors")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from") // TODO: Add output flag
	cmd.Flags().BoolVar(&offline, "offline", false, "Lists offline collectors")
	return cmd
}

func listCollectors(filter string, limit string, offset string, output string, offline bool, log zerolog.Logger) {
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
	if limit != "" && offline == false {
		// TODO: Add validation that string is a number
		query.Add("limit", limit)
	}
	if offset != "" && offline == false {
		// TODO: Validate something?
		query.Add("offset", offset)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
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
		if factory.ValidateCollectorOutput(output) == true {
			value := gjson.Get(string(collectorInfoJson), "#."+output)
			formattedValue := strings.Trim(value.String(), `[]`)
			formattedValue = strings.ReplaceAll(formattedValue, "\"", "")
			fmt.Println(formattedValue)
		} else {
			fmt.Print(string(collectorInfoJson))
		}
	}
}
