package authorizers

import (
	"fmt"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/wizedkyle/sumocli/internal/config"
)

type OAuthGrantType int

var armAuthorizer autorest.Authorizer

const (
	OAuthGrantTypeServicePrincipal OAuthGrantType = iota
	OAuthGrantTypeDeviceFlow
)

func GetARMAuthorizer() (autorest.Authorizer, error) {
	if armAuthorizer != nil {
		return armAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(grantType(), config.Environment().ResourceManagerEndpoint)
	if err == nil {
		armAuthorizer = a
	} else {
		armAuthorizer = nil
	}
	return armAuthorizer, err
}

func getAuthorizerForResource(grantType OAuthGrantType, resource string) (autorest.Authorizer, error) {
	var a autorest.Authorizer
	var err error

	switch grantType {
	case OAuthGrantTypeServicePrincipal:
		oauthConfig, err := adal.NewOAuthConfig(
			config.Environment().ActiveDirectoryEndpoint, config.GetTenantId())
		if err != nil {
			return nil, err
		}

		token, err := adal.NewServicePrincipalToken(
			*oauthConfig, config.GetClientId(), config.GetClientSecret(), resource)
		if err != nil {
			return nil, err
		}
		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeDeviceFlow:
		deviceConfig := auth.NewDeviceFlowConfig(config.GetClientId(), config.GetTenantId())
		deviceConfig.Resource = resource
		a, err = deviceConfig.Authorizer()
		if err != nil {
			return nil, err
		}

	default:
		return a, fmt.Errorf("invalid grant type specified")
	}

	return a, err
}

func grantType() OAuthGrantType {
	if config.GetUseDeviceFlow() {
		return OAuthGrantTypeDeviceFlow
	}
	return OAuthGrantTypeServicePrincipal
}
