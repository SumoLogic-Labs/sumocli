package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"strconv"
	"strings"
)

func NewCmdIngestBudgetsUpdate() *cobra.Command {
	var (
		action         string
		auditThreshold int
		capacityBytes  int
		description    string
		fieldValue     string
		id             string
		merge          bool
		name           string
		resetTime      string
		timezone       string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing ingest budget.",
		Run: func(cmd *cobra.Command, args []string) {
			updateIngestBudget(action, auditThreshold, capacityBytes, description, fieldValue, id, merge,
				name, resetTime, timezone)
		},
	}
	cmd.Flags().StringVar(&action, "action", "", "Specify an action to take when ingest budget's capacity is reached."+
		"Supported values are either stopCollecting or keepCollecting.")
	cmd.Flags().IntVar(&auditThreshold, "auditThreshold", 1, "Specify a percentage of when an ingest budget's capacity usage is logged in the Audit Index")
	cmd.Flags().IntVar(&capacityBytes, "capacityBytes", 0, "Specify the capacity of the ingest budget in bytes.")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the ingest budget")
	cmd.Flags().StringVar(&fieldValue, "fieldValue", "", "Specify the custom field value that is used to assign Collectors to the ingest budget")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.Flags().BoolVar(&merge, "merge", true, "If set to false it will overwrite the ingest budget configuration")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the ingest budget")
	cmd.Flags().StringVar(&resetTime, "resetTime", "", "Specify the reset time of the ingest bidget in HH:MM format")
	cmd.Flags().StringVar(&timezone, "timezone", "", "Specify the timezone of the reset time in IANA Time Zone format")
	cmd.MarkFlagRequired("action")
	cmd.MarkFlagRequired("capacityBytes")
	cmd.MarkFlagRequired("fieldValue")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("resetTime")
	cmd.MarkFlagRequired("timezone")
	return cmd
}

func updateIngestBudget(action string, auditThreshold int, capacityBytes int, description string, fieldValue string, id string, merge bool,
	name string, resetTime string, timezone string) {
	var ingestBudgetResponse api.GetIngestBudget
	log := logging.GetConsoleLogger()
	if merge == true {
		requestUrl := "/v1/ingestBudgets/" + id
		client, request := factory.NewHttpRequest("GET", requestUrl)
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request " + requestUrl)
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}
		err = json.Unmarshal(responseBody, &ingestBudgetResponse)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		if response.StatusCode != 200 {
			log.Fatal().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
		}

		// Building body payload to update the ingest budget based on the differences
		// between the current ingest budget and the desired settings
		requestBodySchema := &api.CreateIngestBudgetRequest{}
		if strings.EqualFold(ingestBudgetResponse.Action, action) {
			requestBodySchema.Action = ingestBudgetResponse.Action
		} else {
			requestBodySchema.Action = action
		}

		if ingestBudgetResponse.AuditThreshold == auditThreshold {
			requestBodySchema.AuditThreshold = ingestBudgetResponse.AuditThreshold
		} else {
			requestBodySchema.AuditThreshold = auditThreshold
		}

		if ingestBudgetResponse.CapacityBytes == capacityBytes {
			requestBodySchema.CapacityBytes = ingestBudgetResponse.CapacityBytes
		} else {
			requestBodySchema.CapacityBytes = capacityBytes
		}

		if strings.EqualFold(ingestBudgetResponse.Description, description) {
			requestBodySchema.Description = ingestBudgetResponse.Description
		} else {
			requestBodySchema.Description = description
		}

		if strings.EqualFold(ingestBudgetResponse.FieldValue, fieldValue) {
			requestBodySchema.FieldValue = ingestBudgetResponse.FieldValue
		} else {
			requestBodySchema.FieldValue = fieldValue
		}

		if strings.EqualFold(ingestBudgetResponse.Name, name) {
			requestBodySchema.Name = ingestBudgetResponse.Name
		} else {
			requestBodySchema.Name = name
		}

		if strings.EqualFold(ingestBudgetResponse.ResetTime, resetTime) {
			requestBodySchema.ResetTime = ingestBudgetResponse.ResetTime
		} else {
			requestBodySchema.ResetTime = resetTime
		}

		if strings.EqualFold(ingestBudgetResponse.TimeZone, timezone) {
			requestBodySchema.Timezone = ingestBudgetResponse.TimeZone
		} else {
			requestBodySchema.Timezone = timezone
		}

		requestBody, err := json.Marshal(requestBodySchema)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request = factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
		response, err = client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request " + requestUrl)
		}
		defer response.Body.Close()
		responseBody, err = io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}
		err = json.Unmarshal(responseBody, &ingestBudgetResponse)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		ingestBudgetResponseJson, err := json.MarshalIndent(ingestBudgetResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("error marshalling response body")
		}
		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, log)
		} else {
			fmt.Println(string(ingestBudgetResponseJson))
		}
	} else {
		requestBodySchema := &api.CreateIngestBudgetRequest{
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
		requestUrl := "/v1/ingestBudgets/" + id
		client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request")
		}

		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}
		err = json.Unmarshal(responseBody, &ingestBudgetResponse)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		ingestBudgetResponseJson, err := json.MarshalIndent(ingestBudgetResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("error marshalling response body")
		}
		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, log)
		} else {
			fmt.Println(string(ingestBudgetResponseJson))
		}
	}
}
