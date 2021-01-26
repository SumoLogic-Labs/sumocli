package config

import (
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/internal/build"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

var (
	clientId        string
	clientSecret    string
	tenantId        string
	subscriptionId  string
	defaultLocation string
	cloudName       = "AzurePublicCloud"
	useDeviceFlow   bool
	userAgent       = "Sumocli " + build.Version
	environment     *azure.Environment
)

func AddAzureFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&clientId, "clientId", "", "Specify the client ID of the Azure AD Application")
	cmd.PersistentFlags().StringVar(&clientSecret, "clientSecret", "", "Specify the client secret of the Azure AD Application")
	cmd.PersistentFlags().StringVar(&tenantId, "tenantId", "", "Specify the tenant ID of the Azure AD tenant")
	cmd.PersistentFlags().StringVar(&subscriptionId, "subscriptionId", "", "Specify the subscription ID of the Azure subscription you want to deploy to")
	cmd.PersistentFlags().StringVar(&defaultLocation, "location", "", "Specify the Azure location to deploy resources to")
	cmd.PersistentFlags().BoolVar(&useDeviceFlow, "useDeviceFlow", false, "Uses device flow authentication, requires clientId and tenantId")
}

func GetClientId() string {
	return clientId
}

func GetClientSecret() string {
	return clientSecret
}

func GetTenantId() string {
	return tenantId
}

func GetSubscriptionId() string {
	return subscriptionId
}

func GetDefaultLocation() string {
	return defaultLocation
}

func GetUseDeviceFlow() bool {
	return useDeviceFlow
}

func GetUserAgent() string {
	return userAgent
}

func Environment() *azure.Environment {
	log := logging.GetConsoleLogger()
	if environment != nil {
		return environment
	}
	env, err := azure.EnvironmentFromName(cloudName)
	if err != nil {
		log.Error().Err(err).Msg("unable to retrieve Azure environment name")
	}
	environment = &env
	return environment
}

func GetAzureLogTags() map[string]*string {
	logTags := map[string]*string{
		"CollectionType": to.StringPtr("Logs"),
		"CreatedBy":      to.StringPtr("Sumocli"),
		"Version":        to.StringPtr(build.Version),
	}
	return logTags
}

func AzureMetricTags() map[string]*string {
	metricTags := map[string]*string{
		"CollectionType": to.StringPtr("Metrics"),
		"CreatedBy":      to.StringPtr("Sumocli"),
		"Version":        to.StringPtr(build.Version),
	}
	return metricTags
}
