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

func NewCmdDynamicParsingList() *cobra.Command {
	var (
		limit int
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all dynamic parsing rules.",
		Run: func(cmd *cobra.Command, args []string) {
			listDynamicParsingRules(limit)
		},
	}
	cmd.Flags().IntVar(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	return cmd
}

func listDynamicParsingRules(limit int) {
	var dynamicParsingRulesResponse api.ListDynamicParsingRules
	log := logging.GetConsoleLogger()
	requestUrl := "v1/dynamicParsingRules"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
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

	err = json.Unmarshal(responseBody, &dynamicParsingRulesResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	dynamicParsingRulesResponseJson, err := json.MarshalIndent(dynamicParsingRulesResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(dynamicParsingRulesResponseJson))
	}
}
