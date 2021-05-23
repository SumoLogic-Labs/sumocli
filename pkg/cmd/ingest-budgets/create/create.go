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

func NewCmdIngestBudgetsCreate() *cobra.Command {
	var (
		action         string
		auditThreshold int
		capacityBytes  int
		description    string
		fieldValue     string
		name           string
		resetTime      string
		timezone       string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new ingest budget.",
		Run: func(cmd *cobra.Command, args []string) {
			createIngestBudget(action, auditThreshold, capacityBytes, description, fieldValue, name, resetTime, timezone)
		},
	}
	cmd.Flags().StringVar(&action, "action", "", "Specify an action to take when ingest budget's capacity is reached."+
		"Supported values are either stopCollecting or keepCollecting.")
	cmd.Flags().IntVar(&auditThreshold, "auditThreshold", 1, "Specify a percentage of when an ingest budget's capacity usage is logged in the Audit Index")
	cmd.Flags().IntVar(&capacityBytes, "capacityBytes", 0, "Specify the capacity of the ingest budget in bytes.")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the ingest budget")
	cmd.Flags().StringVar(&fieldValue, "fieldValue", "", "Specify the custom field value that is used to assign Collectors to the ingest budget")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the ingest budget")
	cmd.Flags().StringVar(&resetTime, "resetTime", "", "Specify the reset time of the ingest bidget in HH:MM format")
	cmd.Flags().StringVar(&timezone, "timezone", "", "Specify the timezone of the reset time in IANA Time Zone format")
	cmd.MarkFlagRequired("action")
	cmd.MarkFlagRequired("capacityBytes")
	cmd.MarkFlagRequired("fieldValue")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("resetTime")
	cmd.MarkFlagRequired("timezone")
	return cmd
}

func createIngestBudget(action string, auditThreshold int, capacityBytes int, description string,
	fieldValue string, name string, resetTime string, timezone string) {
	var ingestBudgetResponse api.GetIngestBudget
	log := logging.GetConsoleLogger()
	requestBodySchema := api.CreateIngestBudgetRequest{
		Name:           name,
		FieldValue:     fieldValue,
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
	requestUrl := "v1/ingestBudgets"
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
