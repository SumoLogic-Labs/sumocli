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
	egSubClient := eventgrid.NewEventSubscriptionsClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	egSubClient.Authorizer = auth
	egSubClient.AddToUserAgent(config.GetUserAgent())
	return egSubClient
}

func GetEventGridTopicClient() eventgrid.TopicsClient {
	egClient := eventgrid.NewTopicsClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	egClient.Authorizer = auth
	egClient.AddToUserAgent(config.GetUserAgent())
	return egClient
}

func GetEventHubClient() eventhub.EventHubsClient {
	ehClient := eventhub.NewEventHubsClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	ehClient.Authorizer = auth
	ehClient.AddToUserAgent(config.GetUserAgent())
	return ehClient
}

func GetEventHubNamespaceClient() eventhub.NamespacesClient {
	ehClient := eventhub.NewNamespacesClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	ehClient.Authorizer = auth
	ehClient.AddToUserAgent(config.GetUserAgent())
	return ehClient
}

func GetInsightsClient() insights.ComponentsClient {
	insightsClient := insights.NewComponentsClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	insightsClient.Authorizer = auth
	insightsClient.AddToUserAgent(config.GetUserAgent())
	return insightsClient
}

func GetNamespaceClient() servicebus.NamespacesClient {
	nsClient := servicebus.NewNamespacesClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	nsClient.Authorizer = auth
	nsClient.AddToUserAgent(config.GetUserAgent())
	return nsClient
}

func GetQueueClient() servicebus.QueuesClient {
	queueClient := servicebus.NewQueuesClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	queueClient.Authorizer = auth
	queueClient.AddToUserAgent(config.GetUserAgent())
	return queueClient
}

func GetResourceGroupClient() features.ResourceGroupsClient {
	rgClient := features.NewResourceGroupsClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	rgClient.Authorizer = auth
	rgClient.AddToUserAgent(config.GetUserAgent())
	return rgClient
}

func GetStorageClient() storage.AccountsClient {
	sgClient := storage.NewAccountsClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	sgClient.Authorizer = auth
	sgClient.AddToUserAgent(config.GetUserAgent())
	return sgClient
}

func GetStorageTableClient() storage.TableClient {
	sgTableClient := storage.NewTableClient(config.GetSubscriptionId())
	auth, _ := authorizers.GetARMAuthorizer()
	sgTableClient.Authorizer = auth
	sgTableClient.AddToUserAgent(config.GetUserAgent())
	return sgTableClient
}
