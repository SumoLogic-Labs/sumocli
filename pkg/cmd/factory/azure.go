package factory

import (
	"fmt"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/wizedkyle/sumocli/internal/build"
	"os"
)

var (
	ApplicationId     string
	ApplicationSecret string
	Cloud             string
	Location          string
	SubscriptionId    string
	TenantId          string
)

func AzureRMAuth() (autorest.Authorizer, error) {
	armAuthorizer, err := getAuthorizerForResource(environment().ResourceManagerEndpoint)
	return armAuthorizer, err
}

func AzureLogTags() map[string]*string {
	logTags := map[string]*string{
		"CollectionType": to.StringPtr("Logs"),
		"CreatedBy":      to.StringPtr("Sumocli"),
		"Version":        to.StringPtr(build.Version),
	}
	return logTags
}

/*
func AzureMetricTags() map[string]string {

}
*/

func environment() *azure.Environment {
	var environment *azure.Environment

	if environment != nil {
		return environment
	}
	env, err := azure.EnvironmentFromName(Cloud)
	if err != nil {
		fmt.Println("Invalid cloud name: " + Cloud + " cannot continue")
		os.Exit(0)
	}
	environment = &env
	return environment
}

func getAuthorizerForResource(resource string) (autorest.Authorizer, error) {
	var auth autorest.Authorizer
	var err error
	if ApplicationId != "" && ApplicationSecret != "" && SubscriptionId != "" && TenantId != "" {
		config, err := adal.NewOAuthConfig(
			environment().ActiveDirectoryEndpoint, TenantId)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		token, err := adal.NewServicePrincipalToken(
			*config, ApplicationId, ApplicationSecret, resource)
		if err != nil {
			return nil, err
		}
		auth := autorest.NewBearerAuthorizer(token)
		return auth, err
	}
	return auth, err
}
