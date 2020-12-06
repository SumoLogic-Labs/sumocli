package create

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/features"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/storage/mgmt/storage"
	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/azure-sdk-for-go/services/eventgrid/mgmt/2020-06-01/eventgrid"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-06-01/web"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"os"
)

func NewCmdAzureCreate() *cobra.Command {
	var (
		prefix     string
		diagnostic bool
		metrics    bool
		blob       bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create Azure infrastructure to collect logs or metrics",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			log := logging.GetConsoleLogger()
			logger.Debug().Msg("Create Azure infrastructure request started")
			if blob == true {
				azureCreateBlobCollection(prefix, log)
			} else if metrics == true {

			} else if diagnostic == true {

			} else {
				fmt.Println("Please select either --diagnostic, --logs or --metrics")
			}
			logger.Debug().Msg("Create Azure infrastructure request finished.")
		},
	}

	cmd.Flags().BoolVar(&blob, "blob", false, "Deploys infrastructure for Azure Blob collection.")
	cmd.Flags().StringVar(&prefix, "prefix", "", "Name of the resource")
	return cmd
}

func azureCreateBlobCollection(prefix string, log zerolog.Logger) {
	ctx := context.Background()
	logsName := "scliblob"
	rgName := logsName + prefix
	sgName := logsName + prefix
	sourceSgName := "sclisrc" + prefix
	nsName := logsName + prefix
	nsAuthName := logsName + prefix
	queueName := logsName + prefix
	ehNsName := logsName + prefix + "ehns"
	ehName := logsName + prefix + "eh"
	ehAuthName := logsName + prefix + "ehrule"
	cgName := logsName + prefix
	topicName := logsName + prefix
	eventSubName := logsName + prefix
	insightsName := logsName + prefix
	appPlanName := logsName + prefix

	createResourceGroup(ctx, rgName, log)
	createStorageAccount(ctx, rgName, sgName, log)
	sourceSgAcc, _ := createStorageAccount(ctx, rgName, sourceSgName, log)
	createStorageAccountTable(ctx, rgName, sgName, log)
	createServiceBusNamespace(ctx, rgName, nsName, log)
	createServiceBusAuthRule(ctx, rgName, sgName, nsAuthName, log)
	createServiceBusQueue(ctx, rgName, nsName, queueName, log)
	createEventHubNamespace(ctx, rgName, ehNsName, log)
	eh := createEventHub(ctx, rgName, ehNsName, ehName, log)
	createEventHubAuthRule(ctx, rgName, ehNsName, ehName, ehAuthName, log)
	createEventHubConsumerGroup(ctx, rgName, ehNsName, ehName, cgName, log)
	createEventGridTopic(ctx, rgName, topicName, log)
	createEventGridSubscription(ctx, to.String(sourceSgAcc.ID), eventSubName, eh, log)
	createApplicationInsight(ctx, rgName, insightsName, log)
	createAppServicePlan(ctx, rgName, appPlanName, log)
}

/*
func azureCreateLogCollection() {
}

func azureCreateMetricCollection() {
}
*/

func createAppServicePlan(ctx context.Context, rgName string, appPlanName string, log zerolog.Logger) (web.AppServicePlan, error) {
	log.Info().Msg("Creating or updating App Service Plan " + appPlanName)
	appClient := factory.GetAppServicePlanClient()
	appPlan, err := appClient.CreateOrUpdate(
		ctx,
		rgName,
		appPlanName,
		web.AppServicePlan{
			AppServicePlanProperties: nil,
			Sku: &web.SkuDescription{
				Name: to.StringPtr("Y1"),
				Tier: to.StringPtr("Dynamic"),
				Size: to.StringPtr("Y1"),
			},
			Kind:     to.StringPtr("FunctionApp"),
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update App Service Plan " + appPlanName)
		os.Exit(0)
	}

	err = appPlan.WaitForCompletionRef(ctx, appClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update App Service Plan " + appPlanName)
	}

	log.Info().Msg("Created or updated App Service Plan " + appPlanName)
	return appPlan.Result(appClient)
}

func createEventGridSubscription(ctx context.Context, scope string, eventSubName string, eventhub eventhub.Model, log zerolog.Logger) eventgrid.EventSubscriptionsCreateOrUpdateFuture {
	log.Info().Msg("Creating or updating Event Grid Subscription " + eventSubName)
	egSubClient := factory.GetEventGridSubscriptionClient()
	subscription, err := egSubClient.CreateOrUpdate(
		ctx,
		scope,
		eventSubName,
		eventgrid.EventSubscription{
			EventSubscriptionProperties: &eventgrid.EventSubscriptionProperties{
				Destination: eventgrid.BasicEventSubscriptionDestination(
					eventgrid.EventHubEventSubscriptionDestination{
						EventHubEventSubscriptionDestinationProperties: &eventgrid.EventHubEventSubscriptionDestinationProperties{
							ResourceID: eventhub.ID,
						},
						EndpointType: eventgrid.EndpointTypeEventHub,
					}),
				Filter: &eventgrid.EventSubscriptionFilter{
					IncludedEventTypes: &[]string{
						"Microsoft.Storage.BlobCreated",
					},
				},
			},
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Event Grid Subscription " + eventSubName)
		os.Exit(0)
	}
	err = subscription.WaitForCompletionRef(ctx, egSubClient.Client)

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Event Grid Subscription " + eventSubName)
		os.Exit(0)
	}

	log.Info().Msg("Created or updated Event Grid Subscription " + eventSubName)
	return subscription
}

func createEventGridTopic(ctx context.Context, rgName string, topicName string, log zerolog.Logger) (eventgrid.Topic, error) {
	log.Info().Msg("Creating or updating Event Grid Topic " + topicName)
	topicClient := factory.GetEventGridTopicClient()
	topic, err := topicClient.CreateOrUpdate(
		ctx,
		rgName,
		topicName,
		eventgrid.Topic{
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Event Grid Topic " + topicName)
		os.Exit(0)
	}

	err = topic.WaitForCompletionRef(ctx, topicClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Event Grid Topic " + topicName)
		os.Exit(0)
	}

	log.Info().Msg("Created or updated Event Grid Topic " + topicName)
	return topic.Result(topicClient)
}

func createEventHubNamespace(ctx context.Context, rgName string, nsName string, log zerolog.Logger) (eventhub.EHNamespace, error) {
	log.Info().Msg("Creating or updating Event Hub namespace " + nsName)
	ehClient := factory.GetEventHubNamespaceClient()
	ehNamespace, err := ehClient.CreateOrUpdate(
		ctx,
		rgName,
		nsName,
		eventhub.EHNamespace{
			Sku: &eventhub.Sku{
				Name:     eventhub.Standard,
				Capacity: to.Int32Ptr(1),
			},
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Event Hub namespace " + nsName)
		os.Exit(0)
	}

	err = ehNamespace.WaitForCompletionRef(ctx, ehClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Event Hub namespace " + nsName)
		os.Exit(0)
	}

	log.Info().Msg("Created or updated Event Hub namespace " + nsName)
	return ehNamespace.Result(ehClient)
}

func createEventHub(ctx context.Context, rgName string, ehNsName string, ehName string, log zerolog.Logger) eventhub.Model {
	log.Info().Msg("Creating or updating Event Hub " + ehName)
	ehClient := factory.GetEventHubClient()
	eh, err := ehClient.CreateOrUpdate(
		ctx,
		rgName,
		ehNsName,
		ehName,
		eventhub.Model{
			Properties: &eventhub.Properties{
				MessageRetentionInDays: to.Int64Ptr(7),
				PartitionCount:         to.Int64Ptr(2),
			},
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update EventHub " + ehName)
		os.Exit(0)
	}
	return eh
}

func createEventHubAuthRule(ctx context.Context, rgName string, ehNsName string, ehName string, ehAuthName string, log zerolog.Logger) {
	log.Info().Msg("Creating or updating Event Hub Authorization Rule " + ehAuthName)
	ehClient := factory.GetEventHubClient()
	_, err := ehClient.CreateOrUpdateAuthorizationRule(
		ctx,
		rgName,
		ehNsName,
		ehName,
		ehAuthName,
		eventhub.AuthorizationRule{
			AuthorizationRuleProperties: &eventhub.AuthorizationRuleProperties{
				Rights: &[]eventhub.AccessRights{
					"Listen",
					"Manage",
					"Send",
				}},
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Event Hub Authorization rule " + ehAuthName)
		os.Exit(0)
	}
	log.Info().Msg("Created or updated Event Hub Authorization rule " + ehAuthName)
}

func createEventHubConsumerGroup(ctx context.Context, rgName string, ehNsName string, ehName string, cgName string, log zerolog.Logger) {
	log.Info().Msg("Creating or updating Event Hub Consumer Group " + cgName)
	csClient := factory.GetConsumerGroupsClient()
	_, err := csClient.CreateOrUpdate(
		ctx,
		rgName,
		ehNsName,
		ehName,
		cgName,
		eventhub.ConsumerGroup{
			ConsumerGroupProperties: nil,
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Event Hub Consumer Group " + cgName)
		os.Exit(0)
	}
	log.Info().Msg("Created or updated Event Hub Consumer Group " + cgName)
}

func createApplicationInsight(ctx context.Context, rgName string, insightsName string, log zerolog.Logger) insights.ApplicationInsightsComponent {
	log.Info().Msg("Creating or updating Application Insights: " + insightsName)
	insightsClient := factory.GetInsightsClient()
	insights, err := insightsClient.CreateOrUpdate(
		ctx,
		rgName,
		insightsName,
		insights.ApplicationInsightsComponent{
			Kind: to.StringPtr("web"),
			ApplicationInsightsComponentProperties: &insights.ApplicationInsightsComponentProperties{
				ApplicationID:              nil,
				AppID:                      nil,
				ApplicationType:            "",
				FlowType:                   "",
				RequestSource:              "",
				InstrumentationKey:         nil,
				CreationDate:               nil,
				TenantID:                   nil,
				HockeyAppID:                nil,
				HockeyAppToken:             nil,
				ProvisioningState:          nil,
				SamplingPercentage:         nil,
				ConnectionString:           nil,
				RetentionInDays:            nil,
				DisableIPMasking:           nil,
				ImmediatePurgeDataOn30Days: nil,
				PrivateLinkScopedResources: nil,
				IngestionMode:              "",
			},
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Application Insights: " + insightsName)
		os.Exit(0)
	}

	log.Info().Msg("Created or updated Application Insights: " + insightsName)
	return insights
}

func createResourceGroup(ctx context.Context, rgName string, log zerolog.Logger) features.ResourceGroup {
	log.Info().Msg("Creating or updating resource group " + rgName)
	rgClient := factory.GetResourceGroupClient()
	rg, err := rgClient.CreateOrUpdate(
		ctx,
		rgName,
		features.ResourceGroup{
			Name:     to.StringPtr(rgName),
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Resource Group " + rgName)
		os.Exit(0)
	}
	log.Info().Msg("Created or updated resource group " + rgName)
	return rg
}

func createStorageAccount(ctx context.Context, rgName string, sgName string, log zerolog.Logger) (storage.Account, error) {
	log.Info().Msg("Creating or updating storage account " + sgName)
	sgClient := factory.GetStorageClient()
	sgAccount, err := sgClient.Create(
		ctx,
		rgName,
		sgName,
		storage.AccountCreateParameters{
			Sku: &storage.Sku{
				Name: storage.StandardLRS,
				Tier: storage.Standard,
			},
			Kind:     storage.StorageV2,
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Storage Account " + sgName)
		os.Exit(0)
	}

	err = sgAccount.WaitForCompletionRef(ctx, sgClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update Storage Account " + sgName)
		os.Exit(0)
	}

	log.Info().Msg("Created or updated Storage Account " + rgName)
	return sgAccount.Result(sgClient)
}

func createStorageAccountTable(ctx context.Context, rgName string, sgName string, log zerolog.Logger) {
	log.Info().Msg("Creating FileOffsetMap table")
	tableClient := factory.GetStorageTableClient()
	_, err := tableClient.Create(
		ctx,
		rgName,
		sgName,
		"FileOffsetMap")

	if err != nil {
		log.Error().Err(err).Msg("cannot create FileOffsetMap table")
		os.Exit(0)
	}

	log.Info().Msg("Created FileOffsetMap table")
}

func createServiceBusNamespace(ctx context.Context, rgName string, nsName string, log zerolog.Logger) (servicebus.SBNamespace, error) {
	log.Info().Msg("Creating or updating Service Bus namespace " + nsName)
	nsClient := factory.GetNamespaceClient()
	ns, err := nsClient.CreateOrUpdate(
		ctx,
		rgName,
		nsName,
		servicebus.SBNamespace{
			Sku: &servicebus.SBSku{
				Name: servicebus.Standard,
			},
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create Service Bus namespace " + nsName)
		os.Exit(0)
	}

	err = ns.WaitForCompletionRef(ctx, nsClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create Service Bus namespace " + nsName)
		os.Exit(0)
	}

	log.Info().Msg("Created or updated Service Bus namespace " + nsName)
	return ns.Result(nsClient)
}

func createServiceBusAuthRule(ctx context.Context, rgName string, nsName string, nsAuthName string, log zerolog.Logger) {
	log.Info().Msg("Creating or updating Service Bus namespace authorization rule " + nsAuthName)
	nsClient := factory.GetNamespaceClient()
	_, err := nsClient.CreateOrUpdateAuthorizationRule(
		ctx,
		rgName,
		nsName,
		nsAuthName,
		servicebus.SBAuthorizationRule{
			SBAuthorizationRuleProperties: &servicebus.SBAuthorizationRuleProperties{
				Rights: &[]servicebus.AccessRights{
					"Listen",
					"Manage",
					"Send",
				},
			},
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create Service Bus namespace authorization rule " + nsAuthName)
		os.Exit(0)
	}
	log.Info().Msg("Created or updated Service Bus namespace authorization rule " + nsAuthName)
}

func createServiceBusQueue(ctx context.Context, rgName string, nsName string, queueName string, log zerolog.Logger) {
	log.Info().Msg("Creating or updating Service Bus Queue " + queueName)
	queueClient := factory.GetQueueClient()
	_, err := queueClient.CreateOrUpdate(
		ctx,
		rgName,
		nsName,
		queueName,
		servicebus.SBQueue{
			SBQueueProperties: &servicebus.SBQueueProperties{
				LockDuration:                        to.StringPtr("PT5M"),
				MaxSizeInMegabytes:                  to.Int32Ptr(2048),
				RequiresDuplicateDetection:          to.BoolPtr(false),
				RequiresSession:                     to.BoolPtr(false),
				DefaultMessageTimeToLive:            to.StringPtr("P14D"),
				DeadLetteringOnMessageExpiration:    to.BoolPtr(true),
				DuplicateDetectionHistoryTimeWindow: to.StringPtr("PT10M"),
				MaxDeliveryCount:                    to.Int32Ptr(10),
				EnableBatchedOperations:             to.BoolPtr(true),
				AutoDeleteOnIdle:                    to.StringPtr("P10675199DT2H48M5.4775807S"),
				EnablePartitioning:                  to.BoolPtr(true),
				EnableExpress:                       to.BoolPtr(true),
			},
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create Service Bus Queue " + queueName)
		os.Exit(0)
	}
	log.Info().Msg("Created or updatd Service Bus Queue " + queueName)
}
