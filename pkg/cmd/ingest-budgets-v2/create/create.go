package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdIngestBudgetsV2Create() *cobra.Command {
	var (
		action         string
		auditThreshold int
		capacityBytes  int
		description    string
		scope          string
		name           string
		resetTime      string
		timezone       string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new ingest budget.",
		Run: func(cmd *cobra.Command, args []string) {
			createIngestBudgetV2(action, auditThreshold, capacityBytes, description, name, resetTime, scope, timezone)
		},
	}
	cmd.Flags().StringVar(&action, "action", "", "Specify an action to take when ingest budget's capacity is reached."+
		"Supported values are either stopCollecting or keepCollecting.")
	cmd.Flags().IntVar(&auditThreshold, "auditThreshold", 1, "Specify a percentage of when an ingest budget's capacity usage is logged in the Audit Index")
	cmd.Flags().IntVar(&capacityBytes, "capacityBytes", 0, "Specify the capacity of the ingest budget in bytes.")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the ingest budget")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the ingest budget")
	cmd.Flags().StringVar(&resetTime, "resetTime", "", "Specify the reset time of the ingest bidget in HH:MM format")
	cmd.Flags().StringVar(&scope, "scope", "", "Specify a scope which will be used to identify the messages on which the budget needs to be applied")
	cmd.Flags().StringVar(&timezone, "timezone", "", "Specify the timezone of the reset time in IANA Time Zone format")
	cmd.MarkFlagRequired("action")
	cmd.MarkFlagRequired("capacityBytes")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("resetTime")
	cmd.MarkFlagRequired("scope")
	cmd.MarkFlagRequired("timezone")
	return cmd
}

func createIngestBudgetV2(action string, auditThreshold int, capacityBytes int, description string, name string,
	resetTime string, scope string, timezone string) {
	var ingestBudgetResponse api.GetIngestBudgetV2
	log := logging.GetConsoleLogger()
	requestBodySchema := api.CreateIngestBudgetV2Request{
		Name:           name,
		Scope:          scope,
		CapacityBytes:  capacityBytes,
		Timezone:       timezone,
		ResetTime:      resetTime,
		Description:    description,
		Action:         action,
		AuditThreshold: auditThreshold,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal request body")
	}
	requestUrl := "/v2/ingestBudgets"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
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
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(ingestBudgetResponseJson))
	}
}
