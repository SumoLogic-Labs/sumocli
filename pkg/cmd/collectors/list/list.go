package list

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/tidwall/gjson"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"io/ioutil"
	"net/url"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"strings"
)



func NewCmdCollectorsList() *cobra.Command {
	var (
		numberOfResults string
		filter          string
		output          string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic collectors",
		Long: `The following fields can be exported using the --output command:
id
name
description
collectorType
alive
collectorVersion
`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Collectors list request started.")
			collectors(numberOfResults, filter, output, logger)
			logger.Debug().Msg("Collectors list request finished.")
		},
	}

	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func collectors(numberOfResults string, name string, output string, logger zerolog.Logger) {
	var collectorList api.CollectorListResponse
	client, request := factory.NewHttpRequest("GET", "v1/collectors")
	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if name != "" {
		query.Add("name", name)
	}
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	logging.LogError(err, logger)

	jsonErr := json.Unmarshal(responseBody, &collectorList)
	logging.LogError(jsonErr, logger)

	collectorsJson, err := json.MarshalIndent(collectorList.Collectors, "", "    ")
	logging.LogError(err, logger)

	// Determines if the response should be written to a file or to console
	if validateOutput(output) == true {
		value := gjson.Get(string(collectorsJson), "#."+output)
		formattedValue := strings.Trim(value.String(), `"[]"`)
		fmt.Println(formattedValue)
	} else {
		fmt.Println(string(collectorsJson))
	}
}

func validateOutput(output string) bool {
	switch output {
	case
		"id",
		"name",
		"description",
		"collectorType",
		"alive",
		"collectorVersion":
		return true
	}
	return false
}