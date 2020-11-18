package get

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
"os"
"strings"
)

func NewCmdCollectorGet() *cobra.Command {
	var (
		id	   string
		name   string
		output string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic collector",
		Long: `The following fields can be exported using the --output command:
name
description
filterPredicate
users
capabilities
id
`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("CollectorListResponse get request started.")
			if id == "" {
				if name != "" {
					getCollectorByName(name, output, logger)
				}
			} else {
				getCollector(id, output, logger)
			}
			logger.Debug().Msg("CollectorListResponse get request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the collector to get")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the collector to get")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func getCollectorByName(name string, output string, logger zerolog.Logger) {
	var collectorResource api.CollectorResponse

	requestUrl := "v1/collectors/name/" + name
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	logging.LogError(err, logger)

	jsonErr := json.Unmarshal(responseBody, &collectorResource)
	logging.LogError(jsonErr, logger)

	collectorInfoJson, err := json.MarshalIndent(collectorResource, "", "    ")
	logging.LogError(err, logger)

	if validateOutput(output) == true {
		value := gjson.Get(string(collectorInfoJson), output)
		formattedValue := strings.Trim(value.String(), `"[]"`)
		fmt.Println(formattedValue)
	} else {
		fmt.Println(string(collectorInfoJson))
	}
}

func getCollector(id string, output string, logger zerolog.Logger) {
	var collectorResource api.CollectorResponse

	if id == "" {
		fmt.Println("--id field needs to be specified.")
		os.Exit(0)
	}

	requestUrl := "v1/collectors/" + id
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	logging.LogError(err, logger)

	jsonErr := json.Unmarshal(responseBody, &collectorResource)
	logging.LogError(jsonErr, logger)

	collectorInfoJson, err := json.MarshalIndent(collectorResource, "", "    ")
	logging.LogError(err, logger)

	if validateOutput(output) == true {
		value := gjson.Get(string(collectorInfoJson), output)
		formattedValue := strings.Trim(value.String(), `"[]"`)
		fmt.Println(formattedValue)
	} else {
		fmt.Println(string(collectorInfoJson))
	}
}

func validateOutput(output string) bool {
	switch output {
	case
		"name",
		"description",
		"filterPredicate",
		"users",
		"capabilities",
		"id":
		return true
	}
	return false
}
