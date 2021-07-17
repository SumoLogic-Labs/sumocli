package create

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdAzureEventHubSourceCreate(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var (
		authorizationRuleName string
		category              string
		collectorId           string
		description           string
		eventHubKey           string
		eventHubName          string
		fieldNames            []string
		fieldValues           []string
		name                  string
		namespace             string
	)
	cmd := &cobra.Command{
		Use: "create",
		Short: "The Azure Event Hubs Source provides a secure endpoint to receive data from Azure Event Hubs. " +
			"It securely stores the required authentication, scheduling, and state tracking information. " +
			"This source is used to collect activity and resource logs from Azure.",
		Run: func(cmd *cobra.Command, args []string) {
			createEventHubSource(authorizationRuleName, category, collectorId, description, eventHubKey, eventHubName,
				fieldNames, fieldValues, name, namespace, client, log)
		},
	}
	cmd.Flags().StringVar(&authorizationRuleName, "authorizationRuleName", "", "Specify the name of the Event Hub Authorization Rule")
	cmd.Flags().StringVar(&category, "category", "", "Specify the source category for the source")
	cmd.Flags().StringVar(&collectorId, "collectorId", "", "Specify the collector id to associate the source to")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the source")
	cmd.Flags().StringVar(&eventHubKey, "eventHubKey", "", "Specify either the primary or secondary Event Hub key")
	cmd.Flags().StringVar(&eventHubName, "eventHubName", "", "Specify the name of the Event Hub")
	cmd.Flags().StringSliceVar(&fieldNames, "fieldNames", []string{}, "Specify the names of fields to add to the source "+
		"{names need to be comma separated e.g. field1,field2")
	cmd.Flags().StringSliceVar(&fieldValues, "fieldValues", []string{}, "Specify the values of fields to add to the source "+
		"(values need to be comma separated e.g. value1,value2")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name for the source")
	cmd.Flags().StringVar(&namespace, "namespace", "", "Specify the name of the Event Hub Namespace")
	cmd.MarkFlagRequired("authorizationRuleName")
	cmd.MarkFlagRequired("category")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("eventHubKey")
	cmd.MarkFlagRequired("eventHubName")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("namespace")
	return cmd
}

func createEventHubSource(authorizationRuleName string, category string, collectorId string, description string, eventHubKey string, eventHubName string,
	fieldNames []string, fieldValues []string, name string, namespace string, client *cip.APIClient, log *zerolog.Logger) {
	fields := cmdutils.GenerateFieldsMap(fieldNames, fieldValues)
	body := types.CreateEventHubSourceRequest{
		ApiVersion: "v1",
		Source: types.EventHubSource{
			SchemaRef: types.EventHubSourceSchema{
				Type: "Azure Event Hubs",
			},
			Config: types.EventHubSourceConfigurationDefinition{
				Name:                    name,
				Description:             description,
				Namespace:               namespace,
				HubName:                 eventHubName,
				AccessPolicyName:        authorizationRuleName,
				AccessPolicyKey:         eventHubKey,
				ConsumerGroup:           "$Default",
				Fields:                  fields,
				Category:                category,
				ReceiveWithLatestOffset: true,
			},
			SourceType: "Universal",
		},
	}
	apiResponse, httpResponse, errorResponse := client.CreateEventHubSource(body, collectorId)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to create azure event hub source")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
