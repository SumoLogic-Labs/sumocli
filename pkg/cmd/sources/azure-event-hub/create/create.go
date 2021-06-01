package create

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/features"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/internal/clients"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"strconv"
	"strings"
)

func NewCmdAzureEventHubSourceCreate() *cobra.Command {
	var (
		billing                string
		capacity               int32
		collectorId            int
		description            string
		fieldNames             string
		fieldValues            string
		location               string
		messageRetentionInDays int64
		name                   string
		partitionCount         int64
	)

	cmd := &cobra.Command{
		Use: "create",
		Short: "The Azure Event Hubs Source provides a secure endpoint to receive data from Azure Event Hubs. " +
			"It securely stores the required authentication, scheduling, and state tracking information. " +
			"This source is used to collect activity and resource logs from Azure.",
		Run: func(cmd *cobra.Command, args []string) {
			createEventHubSource(billing, capacity, collectorId, description, fieldNames, fieldValues,
				location, messageRetentionInDays, name, partitionCount)
		},
	}
	cmd.Flags().StringVar(&billing, "billing", "", "Specify the event hub billing type. "+
		"Accepted values are basic or standard")
	cmd.Flags().Int32Var(&capacity, "capacity", 1, "Specify the event hub throughput limits. "+
		"Accepted values are between 0 and 20")
	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Specify the collector id to associate the source to")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the source")
	cmd.Flags().StringVar(&fieldNames, "fieldNames", "", "Specify the names of fields to add to the source "+
		"{names need to be comma separated e.g. field1,field2")
	cmd.Flags().StringVar(&fieldValues, "fieldValues", "", "Specify the values of fields to add to the source "+
		"(values need to be comma separated e.g. value1,value2")
	cmd.Flags().StringVar(&location, "location", "", "Specify the Azure location to deploy resources to")
	cmd.Flags().Int64Var(&messageRetentionInDays, "messageRetentionInDays", 1, "Specify the number of days to "+
		"retain the messages in the event hub. Accepted values are between 1 and 7 days")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name for the source")
	cmd.Flags().Int64Var(&partitionCount, "partitionCount", 2, "Specify the number of partitions for the "+
		"event hub. Accepted values are between 1 and 32")
	cmd.MarkFlagRequired("billing")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("location")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createEventHubSource(billing string, capacity int32, collectorId int, description string, fieldNames string,
	fieldValues string, location string, messageRetentionInDays int64, name string, partitionCount int64) {
	// TODO: var sourceResponse api.AzureEventHubResponse
	resourceName := "sumocli-eventhub-" + name
	log := logging.GetConsoleLogger()
	resourceGroupClient := clients.GetResourceGroupClient()
	resourceGroup, err := resourceGroupClient.CreateOrUpdate(
		context.TODO(),
		resourceName,
		features.ResourceGroup{
			Name:     to.StringPtr(resourceName),
			Location: to.StringPtr(location),
			Tags:     nil,
		})
	if err != nil {
		log.Error().Err(err).Msg("failed to create or update resource group " + resourceName)
	}

	eventHubNamespaceClient := clients.GetEventHubNamespaceClient()
	_, err = eventHubNamespaceClient.CreateOrUpdate(
		context.TODO(),
		*resourceGroup.Name,
		resourceName,
		eventhub.EHNamespace{
			Sku:      getEventHubBilling(billing, capacity),
			Location: resourceGroup.Location,
			Tags:     nil,
		})
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update event hub namespace " + resourceName)
	}

	eventHubClient := clients.GetEventHubClient()
	eventHub, err := eventHubClient.CreateOrUpdate(
		context.TODO(),
		*resourceGroup.Name,
		resourceName,
		resourceName,
		eventhub.Model{
			Properties: &eventhub.Properties{
				MessageRetentionInDays: to.Int64Ptr(messageRetentionInDays),
				PartitionCount:         to.Int64Ptr(partitionCount),
			},
		})
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update event hub " + resourceName)
	}
	eventHubAuthRule, err := eventHubClient.CreateOrUpdateAuthorizationRule(
		context.TODO(),
		*resourceGroup.Name,
		resourceName,
		resourceName,
		resourceName,
		eventhub.AuthorizationRule{
			AuthorizationRuleProperties: &eventhub.AuthorizationRuleProperties{
				Rights: &[]eventhub.AccessRights{
					"Listen",
				},
			},
		})
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update event hub authorization rule " + resourceName)
	}
	eventHubKey, err := eventHubClient.ListKeys(
		context.TODO(),
		*resourceGroup.Name,
		resourceName,
		*eventHub.Name,
		*eventHubAuthRule.Name)
	if err != nil {
		log.Error().Err(err).Msg("cannot get event hub keys for " + resourceName)
	}

	fieldsMap := make(map[string]string)
	if fieldNames != "" && fieldValues != "" {
		fieldNamesSlice := strings.Split(fieldNames, ",")
		fieldValuesSlice := strings.Split(fieldValues, ",")
		for i, _ := range fieldNamesSlice {
			fieldsMap[fieldNamesSlice[i]] = fieldValuesSlice[i]
			i++
		}
	}
	requestBodySchema := &api.EventHubCollection{
		ApiVersion: "v1",
		Source: api.EventHubSource{
			SchemaRef: api.EventHubSourceSchema{
				Type: "Azure Event Hubs",
			},
		},
		Config: api.EventHubConfig{
			Name:                    name,
			Description:             description,
			Namespace:               resourceName,
			HubName:                 *eventHub.Name,
			AccessPolicyName:        *eventHubAuthRule.Name,
			AccessPolicyKey:         *eventHubKey.PrimaryKey,
			ConsumerGroup:           "$Default",
			Fields:                  fieldsMap,
			ReceiveWithLatestOffset: true,
		},
		SourceType: "Universal",
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal request body")
	}
	requestUrl := "v1/collectors/" + strconv.Itoa(collectorId) + "/sources"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}
	fmt.Println(string(responseBody))
}

func getEventHubBilling(billing string, capacity int32) *eventhub.Sku {
	eventHubSku := &eventhub.Sku{}
	eventHubSku.Capacity = to.Int32Ptr(capacity)
	if billing == "basic" {
		eventHubSku.Name = eventhub.Basic
	} else {
		eventHubSku.Name = eventhub.Standard
	}
	return eventHubSku
}
