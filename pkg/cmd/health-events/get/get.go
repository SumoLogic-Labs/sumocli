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

func NewCmdHealthEventsGet() *cobra.Command {
	var (
		collector              bool
		collectorId            string
		collectorName          string
		id                     string
		ingestBudget           bool
		ingestBudgetFieldValue string
		limit                  int
		logsToMetricsRule      bool
		name                   string
		organisation           bool
		scope                  string
		source                 bool
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a list of all the unresolved events in your account that belong to the supplied resource identifiers.",
		Run: func(cmd *cobra.Command, args []string) {
			getCollectorEvents(collector, collectorId, collectorName, id, ingestBudget, ingestBudgetFieldValue, limit,
				logsToMetricsRule, name, organisation, scope, source)
		},
	}
	cmd.Flags().BoolVar(&collector, "collector", false, "Set to true if the resource is a collector")
	cmd.Flags().StringVar(&collectorId, "collectorId", "", "Specify the collector Id (this is only used when the "+
		"source argument is set.)")
	cmd.Flags().StringVar(&collectorName, "collectorName", "", "Specify the collector name (this is only used when the "+
		"source argument is set.)")
	cmd.Flags().StringVar(&id, "id", "", "Specify the unique id of the resource")
	cmd.Flags().BoolVar(&ingestBudget, "ingestBudget", false, "Set to true if the resource is a ingest budget")
	cmd.Flags().StringVar(&ingestBudgetFieldValue, "ingestBudgetFieldValue", "", "Specify the unique field value "+
		"of the ingest budget v1. This will be empty for v2 budgets. (this is only used when the ingestBudget argument is set.)")
	cmd.Flags().IntVar(&limit, "limit", 100, "Specify the number of health events to return")
	cmd.Flags().BoolVar(&logsToMetricsRule, "logsToMetricsRule", false, "Set to true if the resource is a "+
		"logs to metrics rule.")
	cmd.Flags().StringVar(&name, "name", "Unknown", "Specify the name of the resource if required")
	cmd.Flags().BoolVar(&organisation, "organisation", false, "Set to true of the resource is a organisation.")
	cmd.Flags().StringVar(&scope, "scope", "", "Specify the scope of the ingest budget v2. This will be empty "+
		"for v1 budgets. (this is only used when the ingestBudget argument is set.)")
	cmd.Flags().BoolVar(&source, "source", false, "Set to true of the resource is a source")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getCollectorEvents(collector bool, collectorId string, collectorName string, id string, ingestBudget bool,
	ingestBudgetFieldValue string, limit int, logsToMetricsRule bool, name string, organisation bool, scope string, source bool) {
	var healthEventsResponse api.ListHealthEvent
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateHealthEventRequest{}
	requestBodySchemaData := api.HealthEventRequest{
		CollectorId:            collectorId,
		CollectorName:          collectorName,
		Id:                     id,
		IngestBudgetFieldValue: ingestBudgetFieldValue,
		Name:                   name,
		Scope:                  scope,
	}
	if collector == true {
		requestBodySchemaData.Type = "Collector"
	} else if ingestBudget == true {
		requestBodySchemaData.Type = "IngestBudget"
	} else if logsToMetricsRule == true {
		requestBodySchemaData.Type = "LogsToMetricsRule"
	} else if organisation == true {
		requestBodySchemaData.Type = "Organisation"
	} else if source == true {
		requestBodySchemaData.Type = "Source"
	} else {
		log.Error().Msg("please specify one of the following arguments: collector, ingestBudget, logsToMetricsRule, organisation or source")
	}
	requestBodySchema.Data = append(requestBodySchema.Data, requestBodySchemaData)
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal request body")
	}
	requestUrl := "v1/healthEvents/resources"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &healthEventsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	healthEventsResponseJson, err := json.MarshalIndent(healthEventsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(healthEventsResponseJson))
	}
}
