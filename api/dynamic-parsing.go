package api

type CreateDynamicParsingRule struct {
	Name    string `json:"name"`
	Scope   string `json:"scope"`
	Enabled bool   `json:"enabled"`
}

type ListDynamicParsingRules struct {
	Data []DynamicParsingRules `json:"data"`
}

type DynamicParsingRules struct {
	Name       string `json:"name"`
	Scope      string `json:"scope"`
	Enabled    bool   `json:"enabled"`
	CreatedAt  string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
	ModifiedBy string `json:"modifiedBy"`
	Id         string `json:"id"`
}
