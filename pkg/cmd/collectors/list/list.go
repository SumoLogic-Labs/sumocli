package list

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
	cmdUtil "github.com/wizedkyle/sumocli/pkg/cmdutil"
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
			collectors(numberOfResults, filter, output)
		},
	}

	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")
	cmd.Flags().BoolVar(&output, "output", false, "Output results to a file, defaults to false")

	return cmd
}

func collectors(numberOfResults string, name string, output bool) {
	var collector collector
	client := cmdUtil.GetHttpClient()
	authToken, apiEndpoint := login.ReadCredentials()

	request, err := http.NewRequest("GET", apiEndpoint+"v1/collectors", nil)
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	cmdUtil.LogError(err)
	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if name != "" {
		query.Add("name", name)
	}
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	cmdUtil.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	cmdUtil.LogError(err)

	jsonErr := json.Unmarshal(responseBody, &collector)
	cmdUtil.LogError(jsonErr)

	collectorsJson, err := json.MarshalIndent(collector.Collectors, "", "    ")
	cmdUtil.LogError(err)

	// Determines if the response should be written to a file or to console
	if output == true {
		cmdUtil.OutputToFile(collectorsJson)
	} else {
		fmt.Println(string(collectorsJson))
	}
}