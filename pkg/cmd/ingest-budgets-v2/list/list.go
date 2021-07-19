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

func NewCmdIngestBudgetsV2List() *cobra.Command {
	var limit int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all ingest budgets.",
		Run: func(cmd *cobra.Command, args []string) {
			listIngestBudgetsV2(limit)
		},
	}
	cmd.Flags().IntVar(&limit, "limit", 100, "Specify the number of results to return maximum is 1000")
	return cmd
}

func listIngestBudgetsV2(limit int) {
	var ingestBudgetResponse api.ListIngestBudgetsV2
	log := logging.GetConsoleLogger()
	requestUrl := "/v2/ingestBudgets"
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

	err = json.Unmarshal(responseBody, &ingestBudgetResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	ingestBudgetResponseJson, err := json.MarshalIndent(ingestBudgetResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(ingestBudgetResponseJson))
	}
}
