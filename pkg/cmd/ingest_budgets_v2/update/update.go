package update

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgetsV2Update(client *cip.APIClient) *cobra.Command {
	var (
		action         string
		auditThreshold int32
		capacityBytes  int64
		description    string
		id             string
		merge          bool
		name           string
		resetTime      string
		scope          string
		timezone       string
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing ingest budget.",
		Run: func(cmd *cobra.Command, args []string) {
			updateIngestBudgetV2(action, auditThreshold, capacityBytes, description, id, merge, name,
				resetTime, scope, timezone, client)
		},
	}
	cmd.Flags().StringVar(&action, "action", "", "Specify an action to take when ingest budget's capacity is reached."+
		"Supported values are either stopCollecting or keepCollecting.")
	cmd.Flags().Int32Var(&auditThreshold, "auditThreshold", 1, "Specify a percentage of when an ingest budget's capacity usage is logged in the Audit Index")
	cmd.Flags().Int64Var(&capacityBytes, "capacityBytes", 0, "Specify the capacity of the ingest budget in bytes.")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the ingest budget")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the ingest budget")
	cmd.Flags().BoolVar(&merge, "merge", true, "If set to false it will overwrite the ingest budget configuration")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the ingest budget")
	cmd.Flags().StringVar(&resetTime, "resetTime", "", "Specify the reset time of the ingest bidget in HH:MM format")
	cmd.Flags().StringVar(&scope, "scope", "", "Specify a scope which will be used to identify the messages on which the budget needs to be applied")
	cmd.Flags().StringVar(&timezone, "timezone", "", "Specify the timezone of the reset time in IANA Time Zone format")
	cmd.MarkFlagRequired("action")
	cmd.MarkFlagRequired("capacityBytes")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("resetTime")
	cmd.MarkFlagRequired("scope")
	cmd.MarkFlagRequired("timezone")
	return cmd
}

func updateIngestBudgetV2(action string, auditThreshold int32, capacityBytes int64, description string, id string, merge bool,
	name string, resetTime string, scope string, timezone string, client *cip.APIClient) {
	data, response, err := client.UpdateIngestBudgetV2(types.IngestBudgetDefinitionV2{
		Name:           name,
		Scope:          scope,
		CapacityBytes:  capacityBytes,
		Timezone:       timezone,
		ResetTime:      resetTime,
		Description:    description,
		Action:         action,
		AuditThreshold: auditThreshold,
	},
		id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
