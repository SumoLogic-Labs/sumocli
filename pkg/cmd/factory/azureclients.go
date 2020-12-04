package factory

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/features"
)

func GetResourceGroupClient() features.ResourceGroupsClient {
	rgClient := features.NewResourceGroupsClient(SubscriptionId)
	auth, _ := AzureRMAuth()
	rgClient.Authorizer = auth
	rgClient.AddToUserAgent(userAgent())
	return rgClient
}
