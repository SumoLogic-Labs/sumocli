package update

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdIngestBudgetsUpdate(client *cip.APIClient) *cobra.Command {
	var (
		action         string
		auditThreshold int32
		capacityBytes  int64
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
				name, resetTime, timezone, client)
		},
	}
	cmd.Flags().StringVar(&action, "action", "", "Specify an action to take when ingest budget's capacity is reached."+
		"Supported values are either stopCollecting or keepCollecting.")
	cmd.Flags().Int32Var(&auditThreshold, "auditThreshold", 1, "Specify a percentage of when an ingest budget's capacity usage is logged in the Audit Index")
	cmd.Flags().Int64Var(&capacityBytes, "capacityBytes", 0, "Specify the capacity of the ingest budget in bytes.")
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

func updateIngestBudget(action string, auditThreshold int32, capacityBytes int64, description string, fieldValue string, id string, merge bool,
	name string, resetTime string, timezone string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.UpdateIngestBudget(types.IngestBudgetDefinition{
		Name:           name,
		FieldValue:     fieldValue,
		CapacityBytes:  capacityBytes,
		Timezone:       timezone,
		ResetTime:      resetTime,
		Description:    description,
		Action:         action,
		AuditThreshold: auditThreshold,
	},
		id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
