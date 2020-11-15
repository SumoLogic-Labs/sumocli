package list

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
	cmdUtil "github.com/wizedkyle/sumocli/pkg/cmdutil"
	logging "github.com/wizedkyle/sumocli/pkg/logging"
	"github.com/spf13/cobra"
)

type collector struct {
	Collectors []collectorData `json:"collectors"`
}

type linkData struct {
	Rel string `json:"rel"`
	Href string `json:"href"`
}

type collectorData struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CollectorType string `json:"collectorType"`
	Alive bool `json:"alive"`
	Links []linkData `json:"links"`
	CollectorVersion string `json:"collectorVersion"`
	Ephemeral bool `json:"ephemeral"`
	Description string `json:"description"`
	OsName string `json:"osName"`
	OsArch string `json:"osArch"`
	OsVersion string `json:"osVersion"`
	Category string `json:"category"`
}

func NewCmdControllersList() *cobra.Command {
	var (
		numberOfResults string
		filter          string
		output          bool
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic collectors",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Collector list request started.")
			collectors(numberOfResults, filter, output, logger)
			logger.Debug().Msg("Collector list request finished.")
		},
	}

	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")
	cmd.Flags().BoolVar(&output, "output", false, "Output results to a file, defaults to false")

	return cmd
}

func collectors(numberOfResults string, name string, output bool, logger zerolog.Logger) {
	var collector collector
	client := cmdUtil.GetHttpClient()
	authToken, apiEndpoint := login.ReadCredentials()

	request, err := http.NewRequest("GET", apiEndpoint+"v1/collectors", nil)
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	logging.LogError(err, logger)
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

	jsonErr := json.Unmarshal(responseBody, &collector)
	logging.LogError(jsonErr, logger)

	collectorsJson, err := json.MarshalIndent(collector.Collectors, "", "    ")
	logging.LogError(err, logger)

	// Determines if the response should be written to a file or to console
	if output == true {
		cmdUtil.OutputToFile(collectorsJson)
	} else {
		fmt.Println(string(collectorsJson))
	}
}