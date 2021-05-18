package create

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdIngestBudgetsCreate() *cobra.Command {
	var (
		action         string
		auditThreshold int
		capacityBytes  int
		description    string
		fieldValue     string
		name           string
		resetTime      int
		timezone       string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new ingest budget.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}

func createIngestBudget(action string, auditThreshold int, capacityBytes int, description string,
	fieldValue string, name string, resetTime int, timezone string) {
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
	client, request := factory.NewHttpRequestWithBody("GET", requestUrl, requestBody)
}
