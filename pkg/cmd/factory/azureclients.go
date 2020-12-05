package factory

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/features"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/storage/mgmt/storage"
)

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
