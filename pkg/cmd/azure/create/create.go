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
	cgName := logsName + prefix + "cg"
	eventSubName := logsName + prefix + "sub"
	insightsName := logsName + prefix
	appPlanName := logsName + prefix
	functionName := logsName + prefix
	appRepoUrl := "https://github.com/SumoLogic/sumologic-azure-function"
	branch := "master"
	/*
		consumerAppSettings := &[]web.NameValuePair{
			{ Name: to.StringPtr("FUNCTIONS_EXTENSION_VERSION"), Value: to.StringPtr("~1") },
			{ Name: to.StringPtr("Project"), Value: to.StringPtr("BlockBlobReader/target/consumer_build/") },
			{ Name: to.StringPtr("AzureWebJobsStorage"), Value: to.StringPtr("")},
			{ Name: to.StringPtr("APPINSIGHTS_INSTRUMENTATIONKEY"), Value: to.StringPtr("")},
			{ Name: to.StringPtr("SumoLogEndpoint"), Value: to.StringPtr("")}, // TODO: Need to add this
			{ Name: to.StringPtr("TaskQueueConnectionString"), Value: to.StringPtr("")},
			{ Name: to.StringPtr("WEBSITE_NODE_DEFAULT_VERSION"), Value: to.StringPtr("6.5.0")},
			{ Name: to.StringPtr("FUNCTION_APP_EDIT_MODE"), Value: to.StringPtr("readwrite")},
		}
		dlqAppSettings := &[]web.NameValuePair{
			{ Name: to.StringPtr("FUNCTIONS_EXTENSION_VERSION"), Value: to.StringPtr("~1") },
			{ Name: to.StringPtr("Project"), Value: to.StringPtr("BlockBlobReader/target/dlqprocessor_build/") },
			{ Name: to.StringPtr("AzureWebJobsStorage"), Value: to.StringPtr("")},
			{ Name: to.StringPtr("APPINSIGHTS_INSTRUMENTATIONKEY"), Value: to.StringPtr("")},
			{ Name: to.StringPtr("SumoLogEndpoint"), Value: to.StringPtr("")}, // TODO: Need to add this
			{ Name: to.StringPtr("TaskQueueConnectionString"), Value: to.StringPtr("")},
			{ Name: to.StringPtr(" TASKQUEUE_NAME"), Value: to.StringPtr("")},
			{ Name: to.StringPtr("WEBSITE_NODE_DEFAULT_VERSION"), Value: to.StringPtr("6.5.0")},
			{ Name: to.StringPtr("FUNCTION_APP_EDIT_MODE"), Value: to.StringPtr("readwrite")},
		}
	*/

	createResourceGroup(ctx, rgName, log)
	createStorageAccount(ctx, rgName, sgName, log)
	sourceSgAcc, _ := createStorageAccount(ctx, rgName, sourceSgName, log)
	createStorageAccountTable(ctx, rgName, sgName, log)
	createServiceBusNamespace(ctx, rgName, nsName, log)
	createServiceBusAuthRule(ctx, rgName, sgName, nsAuthName, log)
	sbKey := getServiceBusConnectionString(ctx, rgName, nsName, nsAuthName, log)
	createServiceBusQueue(ctx, rgName, nsName, queueName, log)
	createEventHubNamespace(ctx, rgName, ehNsName, log)
	eh := createEventHub(ctx, rgName, ehNsName, ehName, log)
	createEventHubAuthRule(ctx, rgName, ehNsName, ehName, ehAuthName, log)
	ehKey := getEventHubConnectionString(ctx, rgName, ehNsName, ehName, ehAuthName, log)
	createEventHubConsumerGroup(ctx, rgName, ehNsName, ehName, cgName, log)
	createEventGridSubscription(ctx, sourceSgAcc, eventSubName, eh, log)
	appInsights := createApplicationInsight(ctx, rgName, insightsName, log)
	appServicePlan, _ := createAppServicePlan(ctx, rgName, appPlanName, log)

	// Creates each function app, adds source control integration and provides custom App Settings
	// Blob collection requires three apps:  blob reader, consumer, dlq (dead letter queue)
	readerAppSettings := []web.NameValuePair{
		{Name: to.StringPtr("FUNCTIONS_EXTENSION_VERSION"), Value: to.StringPtr("~1")},
		{Name: to.StringPtr("Project"), Value: to.StringPtr("BlockBlobReader/target/producer_build/")},
		{Name: to.StringPtr("AzureWebJobsDashboard"), Value: to.StringPtr(getStorageAccountConnectionString(ctx, rgName, sgName, log))},
		{Name: to.StringPtr("AzureWebJobsStorage"), Value: to.StringPtr(getStorageAccountConnectionString(ctx, rgName, sgName, log))},
		{Name: to.StringPtr("APPINSIGHTS_INSTRUMENTATIONKEY"), Value: appInsights.InstrumentationKey},
		{Name: to.StringPtr("TABLE_NAME"), Value: to.StringPtr("FileOffsetMap")},
		{Name: to.StringPtr("AzureEventHubConnectionString"), Value: ehKey.PrimaryConnectionString},
		{Name: to.StringPtr("TaskQueueConnectionString"), Value: sbKey.PrimaryConnectionString},
		{Name: to.StringPtr("WEBSITE_NODE_DEFAULT_VERSION"), Value: to.StringPtr("6.5.0")},
		{Name: to.StringPtr("FUNCTION_APP_EDIT_MODE"), Value: to.StringPtr("readwrite")},
		{Name: to.StringPtr("WEBSITE_CONTENTAZUREFILECONNECTIONSTRING"), Value: to.StringPtr(getStorageAccountConnectionString(ctx, rgName, sgName, log))},
		{Name: to.StringPtr("WEBSITE_CONTENTSHARE"), Value: to.StringPtr(sgName)},
	}
	readerFunctionName := functionName + "reader"
	createFunctionApp(ctx, rgName, readerFunctionName, appServicePlan, readerAppSettings, log)
	createFunctionAppSourceControl(ctx, rgName, readerFunctionName, appRepoUrl, branch, log)
}

/*
func azureCreateLogCollection() {
}

func azureCreateMetricCollection() {
}
*/

func createApplicationInsight(ctx context.Context, rgName string, insightsName string, log zerolog.Logger) insights.ApplicationInsightsComponent {
	log.Info().Msg("creating or updating application insights: " + insightsName)
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
		log.Error().Err(err).Msg("cannot create or update application insights: " + insightsName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated application insights: " + insightsName)
	return insights
}

func createAppServicePlan(ctx context.Context, rgName string, appPlanName string, log zerolog.Logger) (web.AppServicePlan, error) {
	log.Info().Msg("creating or updating app service plan " + appPlanName)
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
		log.Error().Err(err).Msg("cannot create or update app service plan " + appPlanName)
		os.Exit(0)
	}

	err = appPlan.WaitForCompletionRef(ctx, appClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update app service plan " + appPlanName)
	}

	log.Info().Msg("created or updated app service plan " + appPlanName)
	return appPlan.Result(appClient)
}

func createEventGridSubscription(ctx context.Context, scope storage.Account, eventSubName string, eventhub eventhub.Model, log zerolog.Logger) eventgrid.EventSubscriptionsCreateOrUpdateFuture {
	log.Info().Msg("creating or updating event grid subscription " + eventSubName)
	egSubClient := factory.GetEventGridSubscriptionClient()
	subscription, err := egSubClient.CreateOrUpdate(
		ctx,
		to.String(scope.ID),
		eventSubName,
		eventgrid.EventSubscription{
			EventSubscriptionProperties: &eventgrid.EventSubscriptionProperties{
				Destination: eventgrid.EventHubEventSubscriptionDestination{
					EventHubEventSubscriptionDestinationProperties: &eventgrid.EventHubEventSubscriptionDestinationProperties{
						ResourceID: eventhub.ID,
					},
					EndpointType: eventgrid.EndpointTypeEventHub,
				},
				Filter: &eventgrid.EventSubscriptionFilter{
					IncludedEventTypes: &[]string{
						"Microsoft.Storage.BlobCreated",
					},
				},
			},
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update event grid subscription " + eventSubName)
		os.Exit(0)
	}
	err = subscription.WaitForCompletionRef(ctx, egSubClient.Client)

	if err != nil {
		log.Error().Err(err).Msg("cannot create or update event grid subscription " + eventSubName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated event grid subscription " + eventSubName)
	return subscription
}

func createEventGridTopic(ctx context.Context, rgName string, topicName string, log zerolog.Logger) (eventgrid.Topic, error) {
	log.Info().Msg("creating or updating event grid topic " + topicName)
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
		log.Error().Err(err).Msg("cannot create or update event grid topic " + topicName)
		os.Exit(0)
	}

	err = topic.WaitForCompletionRef(ctx, topicClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update event grid topic " + topicName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated event grid topic " + topicName)
	return topic.Result(topicClient)
}

func createEventHubNamespace(ctx context.Context, rgName string, nsName string, log zerolog.Logger) (eventhub.EHNamespace, error) {
	log.Info().Msg("creating or updating event hub namespace " + nsName)
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
		log.Error().Err(err).Msg("cannot create or update event hub namespace " + nsName)
		os.Exit(0)
	}

	err = ehNamespace.WaitForCompletionRef(ctx, ehClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update event hub namespace " + nsName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated event hub namespace " + nsName)
	return ehNamespace.Result(ehClient)
}

func createEventHub(ctx context.Context, rgName string, ehNsName string, ehName string, log zerolog.Logger) eventhub.Model {
	log.Info().Msg("creating or updating event hub " + ehName)
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
		log.Error().Err(err).Msg("cannot create or update event hub " + ehName)
		os.Exit(0)
	}
	return eh
}

func createEventHubAuthRule(ctx context.Context, rgName string, ehNsName string, ehName string, ehAuthName string, log zerolog.Logger) eventhub.AuthorizationRule {
	log.Info().Msg("creating or updating event hub authorization rule " + ehAuthName)
	ehClient := factory.GetEventHubClient()
	ehAuthRule, err := ehClient.CreateOrUpdateAuthorizationRule(
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
		log.Error().Err(err).Msg("cannot create or update event hub authorization rule " + ehAuthName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated event hub authorization rule " + ehAuthName)
	return ehAuthRule
}

func createEventHubConsumerGroup(ctx context.Context, rgName string, ehNsName string, ehName string, cgName string, log zerolog.Logger) {
	log.Info().Msg("creating or updating event hub consumer group " + cgName)
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
		log.Error().Err(err).Msg("cannot create or update event hub consumer group " + cgName)
		os.Exit(0)
	}
	log.Info().Msg("created or updated event hub consumer group " + cgName)
}

func getEventHubConnectionString(ctx context.Context, rgName string, ehNsName string, ehName string, ehAuthName string, log zerolog.Logger) eventhub.AccessKeys {
	log.Info().Msg("getting event hub keys for " + ehAuthName)
	ehClient := factory.GetEventHubClient()
	ehKey, err := ehClient.ListKeys(
		ctx,
		rgName,
		ehNsName,
		ehName,
		ehAuthName)

	if err != nil {
		log.Error().Err(err).Msg("cannot get event hub keys for " + ehAuthName)
		os.Exit(0)
	}

	log.Info().Msg("obtained event hub keys for " + ehAuthName)
	return ehKey
}

func createFunctionApp(ctx context.Context, rgName string, functionName string, appSerivceId web.AppServicePlan, appSettings []web.NameValuePair, log zerolog.Logger) web.AppsCreateOrUpdateFuture {
	log.Info().Msg("creating or updating azure function " + functionName)
	appClient := factory.GetAppServiceClient()
	functionApp, err := appClient.CreateOrUpdate(
		ctx,
		rgName,
		functionName,
		web.Site{
			SiteProperties: &web.SiteProperties{ // TODO: See if I can add storage account id and key
				Enabled:      to.BoolPtr(true),
				ServerFarmID: appSerivceId.ID,
				SiteConfig: &web.SiteConfig{
					AppSettings: &appSettings,
					ScmType:     web.ScmTypeNone,
				},
				ClientAffinityEnabled: to.BoolPtr(true),
				DailyMemoryTimeQuota:  to.Int32Ptr(1000),
				HTTPSOnly:             to.BoolPtr(true),
			},
			Identity: &web.ManagedServiceIdentity{
				Type: web.ManagedServiceIdentityTypeSystemAssigned,
			},
			Kind:     to.StringPtr("FunctionApp"),
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create azure function app " + functionName)
		os.Exit(0)
	}

	err = functionApp.WaitForCompletionRef(ctx, appClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create azure function app " + functionName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated azure function app " + functionName)
	return functionApp
}

func createFunctionAppSourceControl(ctx context.Context, rgName string, functionName string, appRepoUrl string, branch string, log zerolog.Logger) web.AppsCreateOrUpdateSourceControlFuture {
	log.Info().Msg("creating or updating source control for function app " + functionName)
	appClient := factory.GetAppServiceClient()
	functionAppSc, err := appClient.CreateOrUpdateSourceControl(
		ctx,
		rgName,
		functionName,
		web.SiteSourceControl{
			SiteSourceControlProperties: &web.SiteSourceControlProperties{
				RepoURL:             to.StringPtr(appRepoUrl),
				Branch:              to.StringPtr(branch),
				IsManualIntegration: to.BoolPtr(true),
			},
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create source control settings on function app " + functionName)
		os.Exit(0)
	}

	err = functionAppSc.WaitForCompletionRef(ctx, appClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create source control settings on function app " + functionName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated source control settings on function app " + functionName)
	return functionAppSc
}

func createResourceGroup(ctx context.Context, rgName string, log zerolog.Logger) features.ResourceGroup {
	log.Info().Msg("creating or updating resource group " + rgName)
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
		log.Error().Err(err).Msg("cannot create or update resource group " + rgName)
		os.Exit(0)
	}
	log.Info().Msg("created or updated resource group " + rgName)
	return rg
}

func createStorageAccount(ctx context.Context, rgName string, sgName string, log zerolog.Logger) (storage.Account, error) {
	log.Info().Msg("creating or updating storage account " + sgName)
	sgClient := factory.GetStorageClient()

	// TODO: add storage account name check
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
		log.Error().Err(err).Msg("cannot create or update storage account " + sgName)
		os.Exit(0)
	}

	err = sgAccount.WaitForCompletionRef(ctx, sgClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create or update storage account " + sgName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated storage account " + rgName)
	return sgAccount.Result(sgClient)
}

func createStorageAccountTable(ctx context.Context, rgName string, sgName string, log zerolog.Logger) {
	log.Info().Msg("creating FileOffsetMap table")
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

	log.Info().Msg("created FileOffsetMap table")
}

func getStorageAccountConnectionString(ctx context.Context, rgName string, sgName string, log zerolog.Logger) string {
	log.Info().Msg("getting storage account connection string for " + sgName)
	sgClient := factory.GetStorageClient()
	sgKey, err := sgClient.ListKeys(
		ctx,
		rgName,
		sgName,
		storage.Kerb)

	if err != nil {
		log.Error().Err(err).Msg("cannot get storage account keys")
		os.Exit(0)
	}

	log.Info().Msg("connection string obtained for storage account " + sgName)
	return fmt.Sprintf("DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s;EndpointSuffix=core.windows.net", sgName, to.String((*sgKey.Keys)[0].Value))
}

func createServiceBusNamespace(ctx context.Context, rgName string, nsName string, log zerolog.Logger) (servicebus.SBNamespace, error) {
	log.Info().Msg("creating or updating service bus namespace " + nsName)
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
		log.Error().Err(err).Msg("cannot create service bus namespace " + nsName)
		os.Exit(0)
	}

	err = ns.WaitForCompletionRef(ctx, nsClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create service bus namespace " + nsName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated service bus namespace " + nsName)
	return ns.Result(nsClient)
}

func createServiceBusAuthRule(ctx context.Context, rgName string, nsName string, nsAuthName string, log zerolog.Logger) servicebus.SBAuthorizationRule {
	log.Info().Msg("creating or updating service bus namespace authorization rule " + nsAuthName)
	nsClient := factory.GetNamespaceClient()
	sbAuthRule, err := nsClient.CreateOrUpdateAuthorizationRule(
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
		log.Error().Err(err).Msg("cannot create service bus namespace authorization rule " + nsAuthName)
		os.Exit(0)
	}

	log.Info().Msg("created or updated service bus namespace authorization rule " + nsAuthName)
	return sbAuthRule
}

func getServiceBusConnectionString(ctx context.Context, rgName string, nsName string, nsAuthName string, log zerolog.Logger) servicebus.AccessKeys {
	log.Info().Msg("getting service bus connection string for " + nsAuthName)
	nsClient := factory.GetNamespaceClient()
	sbKeys, err := nsClient.ListKeys(
		ctx,
		rgName,
		nsName,
		nsAuthName)

	if err != nil {
		log.Error().Err(err).Msg("cannot get keys for service bus " + nsAuthName)
		os.Exit(0)
	}

	log.Info().Msg("obtained service bus connection string for " + nsAuthName)
	return sbKeys
}

func createServiceBusQueue(ctx context.Context, rgName string, nsName string, queueName string, log zerolog.Logger) {
	log.Info().Msg("creating or updating service bus queue " + queueName)
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
		log.Error().Err(err).Msg("cannot create service bus queue " + queueName)
		os.Exit(0)
	}
	log.Info().Msg("created or updated service bus queue " + queueName)
}
