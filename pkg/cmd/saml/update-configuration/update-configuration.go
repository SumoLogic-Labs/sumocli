package update_configuration

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/api"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Labs/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
	"reflect"
	"strconv"
	"strings"
)

func NewCmdSamlUpdateConfiguration() *cobra.Command {
	var (
		spInitiatedLoginPath       string
		configurationName          string
		issuer                     string
		spInitiatedLoginEnabled    bool
		authnRequestUrl            string
		x509cert1                  string
		x509cert2                  string
		x509cert3                  string
		firstNameAttribute         string
		lastNameAttribute          string
		onDemandProvisioningRoles  []string
		rolesAttribute             string
		logoutEnabled              bool
		logoutUrl                  string
		emailAttribute             string
		debugMode                  bool
		signAuthnRequest           bool
		disableRequestAuthnContext bool
		isRedirectBinding          bool
		id                         string
		merge                      bool
	)

	cmd := &cobra.Command{
		Use:   "update-configuration",
		Short: "Update an existing SAML configuration in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			updateSamlConfiguration(spInitiatedLoginPath, configurationName, issuer, spInitiatedLoginEnabled,
				authnRequestUrl, x509cert1, x509cert2, x509cert3, firstNameAttribute, lastNameAttribute, onDemandProvisioningRoles,
				rolesAttribute, logoutEnabled, logoutUrl, emailAttribute, debugMode, signAuthnRequest,
				disableRequestAuthnContext, isRedirectBinding, id, merge)
		},
	}
	cmd.Flags().StringVar(&spInitiatedLoginPath, "spInitiatedLoginPath", "", "Specify the identifier used to generate a unique URL for user login")
	cmd.Flags().StringVar(&configurationName, "configurationName", "", "Specify the name of the SSO policy or another name used to described the policy internally")
	cmd.Flags().StringVar(&issuer, "issuer", "", "Specify the unique URL assigned to the organization by the SAML Identity Provider")
	cmd.Flags().BoolVar(&spInitiatedLoginEnabled, "spInitiatedLoginEnabled", false, "Set to true if Sumo Logic redirects users to your identity provider with a SAML AuthnRequest when signing in")
	cmd.Flags().StringVar(&authnRequestUrl, "authnRequestUrl", "", "Specify the URL that the identity provider has assigned for Sumo Logic to submit SAML authentication requests to the identity provider")
	cmd.Flags().StringVar(&x509cert1, "x509cert1", "", "Specify the certificate that is used to verify the signature in SAML assertions")
	cmd.Flags().StringVar(&x509cert2, "x509cert2", "", "Specify a backup certificate used to verify the signature in SAML assertions when x509cert1 expires (optional)")
	cmd.Flags().StringVar(&x509cert3, "x509cert3", "", "Specify a backup certificate used to verify the signature in SAML assertions when x509cert1 expires and x509cert2 is empty (optional)")
	cmd.Flags().StringVar(&firstNameAttribute, "firstNameAttribute", "", "Specify the first name attribute of the new user account")
	cmd.Flags().StringVar(&lastNameAttribute, "lastNameAttribute", "", "Specify the last name attribute of the new user account")
	cmd.Flags().StringSliceVar(&onDemandProvisioningRoles, "onDemandProvisioningRoles", []string{}, "Sumo Logic RBAC roles to be assigned when user accounts are provisioned"+
		"(the roles need to be comma separated e.g. role1,role2,role3)")
	cmd.Flags().StringVar(&rolesAttribute, "rolesAttribute", "", "Specify the role that Sumo Logic will assign to users when they sign in")
	cmd.Flags().BoolVar(&logoutEnabled, "logoutEnabled", false, "Set to true if users are redirected to a URL after signing out of Sumo Logic")
	cmd.Flags().StringVar(&logoutUrl, "logoutUrl", "", "Specify the URL that users will be redirected to after signing out of Sumo Logic")
	cmd.Flags().StringVar(&emailAttribute, "emailAttribute", "", "Specify the email address of the new user account.")
	cmd.Flags().BoolVar(&debugMode, "debugMode", false, "Set to true if additional details are included when a user fails to sign in")
	cmd.Flags().BoolVar(&signAuthnRequest, "signAuthnRequest", false, "Set to true if Sumo Logic will send signed Authn requests to the identity provider")
	cmd.Flags().BoolVar(&disableRequestAuthnContext, "disableRequestAuthnContext", false, "Set to true if Sumo Logic will include the RequestedAuthnContext element of the SAML AuthnRequests it sends to the identity provider")
	cmd.Flags().BoolVar(&isRedirectBinding, "isRedirectBinding", false, "Set to true if the SAML binding is of HTTP Redirect type")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the SAML configuration")
	cmd.Flags().BoolVar(&merge, "merge", true, "If set to false it will overwrite the ingest budget configuration")
	cmd.MarkFlagRequired("configurationName")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("issuer")
	cmd.MarkFlagRequired("x509cert1")
	return cmd
}

func updateSamlConfiguration(spInitiatedLoginPath string, configurationName string, issuer string, spInitiatedLoginEnabled bool,
	authnRequestUrl string, x509cert1 string, x509cert2 string, x509cert3 string, firstNameAttribute string, lastNameAttribute string,
	onDemandProvisioningRoles []string, rolesAttribute string, logoutEnabled bool, logoutUrl string, emailAttribute string,
	debugMode bool, signAuthnRequest bool, disableRequestAuthnContext bool, isRedirectBinding bool, id string, merge bool) {
	var samlResponse api.GetSaml
	log := logging.GetConsoleLogger()
	if merge == true {
		requestUrl := "/v1/saml/identityProviders/" + id
		client, request := factory.NewHttpRequest("GET", requestUrl)
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request " + requestUrl)
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}
		err = json.Unmarshal(responseBody, &samlResponse)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		if response.StatusCode != 200 {
			log.Fatal().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
		}

		// Building body payload to update the saml configuration based on differences
		// between the current saml configuration and the desired settings
		requestBodySchema := &api.CreateSamlRequest{}
		if strings.EqualFold(samlResponse.SpInitiatedLoginPath, spInitiatedLoginPath) {
			requestBodySchema.SpInitiatedLoginPath = samlResponse.SpInitiatedLoginPath
		} else {
			requestBodySchema.SpInitiatedLoginPath = spInitiatedLoginPath
		}

		if strings.EqualFold(samlResponse.ConfigurationName, configurationName) {
			requestBodySchema.ConfigurationName = samlResponse.ConfigurationName
		} else {
			requestBodySchema.ConfigurationName = configurationName
		}

		if strings.EqualFold(samlResponse.Issuer, issuer) {
			requestBodySchema.Issuer = samlResponse.Issuer
		} else {
			requestBodySchema.Issuer = issuer
		}

		if samlResponse.SpInitiatedLoginEnabled == spInitiatedLoginEnabled {
			requestBodySchema.SpInitiatedLoginEnabled = samlResponse.SpInitiatedLoginEnabled
		} else {
			requestBodySchema.SpInitiatedLoginEnabled = spInitiatedLoginEnabled
		}

		if strings.EqualFold(samlResponse.AuthnRequestUrl, authnRequestUrl) {
			requestBodySchema.AuthnRequestUrl = samlResponse.AuthnRequestUrl
		} else {
			requestBodySchema.AuthnRequestUrl = authnRequestUrl
		}

		if strings.EqualFold(samlResponse.X509Cert1, x509cert1) {
			requestBodySchema.X509Cert1 = samlResponse.X509Cert1
		} else {
			requestBodySchema.X509Cert1 = x509cert1
		}

		if strings.EqualFold(samlResponse.X509Cert2, x509cert2) {
			requestBodySchema.X509Cert2 = samlResponse.X509Cert2
		} else {
			requestBodySchema.X509Cert2 = x509cert2
		}

		if strings.EqualFold(samlResponse.X509Cert3, x509cert3) {
			requestBodySchema.X509Cert3 = samlResponse.X509Cert3
		} else {
			requestBodySchema.X509Cert3 = x509cert3
		}

		if strings.EqualFold(samlResponse.OnDemandProvisioningEnabled.FirstNameAttribute, firstNameAttribute) {
			requestBodySchema.OnDemandProvisioningEnabled.FirstNameAttribute = samlResponse.OnDemandProvisioningEnabled.FirstNameAttribute
		} else {
			requestBodySchema.OnDemandProvisioningEnabled.FirstNameAttribute = firstNameAttribute
		}

		if strings.EqualFold(samlResponse.OnDemandProvisioningEnabled.LastNameAttribute, lastNameAttribute) {
			requestBodySchema.OnDemandProvisioningEnabled.LastNameAttribute = samlResponse.OnDemandProvisioningEnabled.LastNameAttribute
		} else {
			requestBodySchema.OnDemandProvisioningEnabled.LastNameAttribute = lastNameAttribute
		}

		if reflect.DeepEqual(samlResponse.OnDemandProvisioningEnabled.OnDemandProvisioningRoles, onDemandProvisioningRoles) {
			requestBodySchema.OnDemandProvisioningEnabled.OnDemandProvisioningRoles = samlResponse.OnDemandProvisioningEnabled.OnDemandProvisioningRoles
		} else {
			requestBodySchema.OnDemandProvisioningEnabled.OnDemandProvisioningRoles = append(requestBodySchema.OnDemandProvisioningEnabled.OnDemandProvisioningRoles, samlResponse.OnDemandProvisioningEnabled.OnDemandProvisioningRoles...)
			requestBodySchema.OnDemandProvisioningEnabled.OnDemandProvisioningRoles = append(requestBodySchema.OnDemandProvisioningEnabled.OnDemandProvisioningRoles, onDemandProvisioningRoles...)
		}

		if strings.EqualFold(samlResponse.RolesAttribute, rolesAttribute) {
			requestBodySchema.RolesAttribute = samlResponse.RolesAttribute
		} else {
			requestBodySchema.RolesAttribute = rolesAttribute
		}

		if samlResponse.LogoutEnabled == logoutEnabled {
			requestBodySchema.LogoutEnabled = samlResponse.LogoutEnabled
		} else {
			requestBodySchema.LogoutEnabled = logoutEnabled
		}

		if strings.EqualFold(samlResponse.LogoutUrl, logoutUrl) {
			requestBodySchema.LogoutUrl = samlResponse.LogoutUrl
		} else {
			requestBodySchema.LogoutUrl = logoutUrl
		}

		if strings.EqualFold(samlResponse.EmailAttribute, emailAttribute) {
			requestBodySchema.EmailAttribute = samlResponse.EmailAttribute
		} else {
			requestBodySchema.EmailAttribute = emailAttribute
		}

		if samlResponse.DebugMode == debugMode {
			requestBodySchema.DebugMode = samlResponse.DebugMode
		} else {
			requestBodySchema.DebugMode = debugMode
		}

		if samlResponse.SignAuthnRequest == signAuthnRequest {
			requestBodySchema.SignAuthnRequest = samlResponse.SignAuthnRequest
		} else {
			requestBodySchema.SignAuthnRequest = signAuthnRequest
		}

		if samlResponse.DisableRequestedAuthnContext == disableRequestAuthnContext {
			requestBodySchema.DisableRequestedAuthnContext = samlResponse.DisableRequestedAuthnContext
		} else {
			requestBodySchema.DisableRequestedAuthnContext = disableRequestAuthnContext
		}

		if samlResponse.IsRedirectBinding == isRedirectBinding {
			requestBodySchema.IsRedirectBinding = samlResponse.IsRedirectBinding
		} else {
			requestBodySchema.IsRedirectBinding = isRedirectBinding
		}

		requestBody, err := json.Marshal(requestBodySchema)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request = factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
		response, err = client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request " + requestUrl)
		}
		defer response.Body.Close()
		responseBody, err = io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("error reading response body from request")
		}
		err = json.Unmarshal(responseBody, &samlResponse)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
		samlResponseJson, err := json.MarshalIndent(samlResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("error marshalling response body")
		}
		if response.StatusCode != 200 {
			log.Error().Msg("Error code = " + strconv.Itoa(response.StatusCode) + string(responseBody))
		} else {
			fmt.Println(string(samlResponseJson))
		}
	} else {
		requestBodySchema := &api.CreateSamlRequest{
			SpInitiatedLoginPath:    spInitiatedLoginPath,
			ConfigurationName:       configurationName,
			Issuer:                  issuer,
			SpInitiatedLoginEnabled: spInitiatedLoginEnabled,
			AuthnRequestUrl:         authnRequestUrl,
			X509Cert1:               x509cert1,
			X509Cert2:               x509cert2,
			X509Cert3:               x509cert3,
			OnDemandProvisioningEnabled: api.OnDemandProvisioningDetail{
				FirstNameAttribute:        firstNameAttribute,
				LastNameAttribute:         lastNameAttribute,
				OnDemandProvisioningRoles: onDemandProvisioningRoles,
			},
			RolesAttribute:               rolesAttribute,
			LogoutEnabled:                logoutEnabled,
			LogoutUrl:                    logoutUrl,
			EmailAttribute:               emailAttribute,
			DebugMode:                    debugMode,
			SignAuthnRequest:             signAuthnRequest,
			DisableRequestedAuthnContext: disableRequestAuthnContext,
			IsRedirectBinding:            isRedirectBinding,
		}
		requestBody, err := json.Marshal(requestBodySchema)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		requestUrl := "/v1/saml/identityProviders/" + id
		client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request")
		}

		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}

		err = json.Unmarshal(responseBody, &samlResponse)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}

		samlResponseJson, err := json.MarshalIndent(samlResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
		}

		if response.StatusCode != 200 {
			factory.HttpError(response.StatusCode, responseBody, log)
		} else {
			fmt.Println(string(samlResponseJson))
		}
	}
}
