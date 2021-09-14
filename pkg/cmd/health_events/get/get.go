package get

import (
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdHealthEventsGet(client *cip.APIClient) *cobra.Command {
	var (
		collector              bool
		collectorId            string
		collectorName          string
		id                     string
		ingestBudget           bool
		ingestBudgetFieldValue string
		limit                  int32
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
				logsToMetricsRule, name, organisation, scope, source, client)
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
	cmd.Flags().Int32Var(&limit, "limit", 100, "Specify the number of health events to return")
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
	ingestBudgetFieldValue string, limit int32, logsToMetricsRule bool, name string, organisation bool, scope string, source bool,
	client *cip.APIClient) {
	var options types.HealthEventsOpts
	var paginationToken string
	options.Limit = optional.NewInt32(limit)
	body := types.ResourceIdentities{}
	bodyData := types.ResourceIdentity{
		CollectorId:            collectorId,
		CollectorName:          collectorName,
		Id:                     id,
		IngestBudgetFieldValue: ingestBudgetFieldValue,
		Name:                   name,
		Scope:                  scope,
	}
	if collector == true {
		bodyData.Type_ = "Collector"
	} else if organisation == true {
		bodyData.Type_ = "Organisation"
	} else if logsToMetricsRule == true {
		bodyData.Type_ = "LogsToMetricsRule"
	} else if ingestBudget == true {
		bodyData.Type_ = "IngestBudget"
	} else if source == true {
		bodyData.Type_ = "Source"
	} else {
		fmt.Println("Please specify one of the following arguments: collector, ingestBudget, logsToMetricsRule, organisation or source")
		os.Exit(1)
	}
	body.Data = append(body.Data, bodyData)
	data, response, err := client.ListAllHealthEventsForResources(body, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	paginationToken = data.Next
	for paginationToken != "" {
		data = listAllHealthEventsForResourcesPagination(client, body, options, paginationToken)
		paginationToken = data.Next
	}
}

func listAllHealthEventsForResourcesPagination(client *cip.APIClient, body types.ResourceIdentities, options types.HealthEventsOpts, token string) types.ListHealthEventResponse {
	options.Token = optional.NewString(token)
	data, response, err := client.ListAllHealthEventsForResources(body, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
	return data
}
