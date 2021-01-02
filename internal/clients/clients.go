package clients

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/features"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/storage/mgmt/storage"
	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/azure-sdk-for-go/services/eventgrid/mgmt/2020-06-01/eventgrid"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-06-01/web"
	"github.com/wizedkyle/sumocli/internal/authorizers"
	"github.com/wizedkyle/sumocli/internal/config"
)

func GetAppServiceClient() web.AppsClient {
	appClient := web.NewAppsClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	appClient.Authorizer = auth
	appClient.AddToUserAgent(config.GetUserAgent())
	return appClient
}

func GetAppServicePlanClient() web.AppServicePlansClient {
	appClient := web.NewAppServicePlansClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	appClient.Authorizer = auth
	appClient.AddToUserAgent(config.GetUserAgent())
	return appClient
}

func GetConsumerGroupsClient() eventhub.ConsumerGroupsClient {
	csClient := eventhub.NewConsumerGroupsClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	csClient.Authorizer = auth
	csClient.AddToUserAgent(config.GetUserAgent())
	return csClient
}

func GetEventGridSubscriptionClient() eventgrid.EventSubscriptionsClient {
	egSubClient := eventgrid.NewEventSubscriptionsClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	egSubClient.Authorizer = auth
	egSubClient.AddToUserAgent(userAgent())
	return egSubClient
}

func GetEventGridTopicClient() eventgrid.TopicsClient {
	egClient := eventgrid.NewTopicsClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	egClient.Authorizer = auth
	egClient.AddToUserAgent(userAgent())
	return egClient
}

func GetEventHubClient() eventhub.EventHubsClient {
	ehClient := eventhub.NewEventHubsClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	ehClient.Authorizer = auth
	ehClient.AddToUserAgent(userAgent())
	return ehClient
}

func GetEventHubNamespaceClient() eventhub.NamespacesClient {
	ehClient := eventhub.NewNamespacesClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	ehClient.Authorizer = auth
	ehClient.AddToUserAgent(userAgent())
	return ehClient
}

func GetInsightsClient() insights.ComponentsClient {
	insightsClient := insights.NewComponentsClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	insightsClient.Authorizer = auth
	insightsClient.AddToUserAgent(userAgent())
	return insightsClient
}

func GetNamespaceClient() servicebus.NamespacesClient {
	nsClient := servicebus.NewNamespacesClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	nsClient.Authorizer = auth
	nsClient.AddToUserAgent(userAgent())
	return nsClient
}

func GetQueueClient() servicebus.QueuesClient {
	queueClient := servicebus.NewQueuesClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	queueClient.Authorizer = auth
	queueClient.AddToUserAgent(userAgent())
	return queueClient
}

func GetResourceGroupClient() features.ResourceGroupsClient {
	rgClient := features.NewResourceGroupsClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	rgClient.Authorizer = auth
	rgClient.AddToUserAgent(userAgent())
	return rgClient
}

func GetStorageClient() storage.AccountsClient {
	sgClient := storage.NewAccountsClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	sgClient.Authorizer = auth
	sgClient.AddToUserAgent(userAgent())
	return sgClient
}

func GetStorageTableClient() storage.TableClient {
	sgTableClient := storage.NewTableClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	sgTableClient.Authorizer = auth
	sgTableClient.AddToUserAgent(userAgent())
	return sgTableClient
}
