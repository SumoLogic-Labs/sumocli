package api

type GetSamlAllowListUsers struct {
	UserId        string `json:"userId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	CanManageSaml bool   `json:"canManageSaml"`
	IsActive      bool   `json:"isActive"`
	LastLogin     string `json:"lastLogin"`
}

type GetSaml struct {
	SpInitiatedLoginPath         string                     `json:"spInitiatedLoginPath"`
	ConfigurationName            string                     `json:"configurationName"`
	Issuer                       string                     `json:"issuer"`
	SpInitiatedLoginEnabled      bool                       `json:"spInitiatedLoginEnabled"`
	AuthnRequestUrl              string                     `json:"authnRequestUrl"`
	X509Cert1                    string                     `json:"x509cert1"`
	X509Cert2                    string                     `json:"x509cert2"`
	X509Cert3                    string                     `json:"x509cert3"`
	OnDemandProvisioningEnabled  onDemandProvisioningDetail `json:"onDemandProvisioningEnabled"`
	RolesAttribute               string                     `json:"rolesAttribute"`
	LogoutEnabled                bool                       `json:"logoutEnabled"`
	LogoutUrl                    string                     `json:"logoutUrl"`
	EmailAttribute               string                     `json:"emailAttribute"`
	DebugMode                    bool                       `json:"debugMode"`
	SignAuthnRequest             bool                       `json:"signAuthnRequest"`
	DisableRequestedAuthnContext bool                       `json:"disableRequestedAuthnContext"`
	IsRedirectBinding            bool                       `json:"isRedirectBinding"`
	Certificate                  string                     `json:"certificate"`
	CreatedAt                    string                     `json:"createdAt"`
	CreatedBy                    string                     `json:"createdBy"`
	ModifiedAt                   string                     `json:"modifiedAt"`
	ModifiedBy                   string                     `json:"modifiedBy"`
	Id                           string                     `json:"id"`
}

type onDemandProvisioningDetail struct {
	FirstNameAttribute        string   `json:"firstNameAttribute"`
	LastNameAttribute         string   `json:"lastNameAttribute"`
	OnDemandProvisioningRoles []string `json:"onDemandProvisioningRoles"`
}
