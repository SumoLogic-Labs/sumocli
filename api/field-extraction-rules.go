package api

type CreateFieldExtractionRule struct {
	Name            string `json:"name"`
	Scope           string `json:"scope"`
	ParseExpression string `json:"parseExpression"`
	Enabled         bool   `json:"enabled"`
}

type ListFieldExtractionRules struct {
	Data []FieldExtractionRules `json:"data"`
}

type FieldExtractionRules struct {
	Name            string   `json:"name"`
	Scope           string   `json:"scope"`
	ParseExpression string   `json:"parseExpression"`
	Enabled         bool     `json:"enabled"`
	CreatedAt       string   `json:"createdAt"`
	CreatedBy       string   `json:"createdBy"`
	ModifiedAt      string   `json:"modifiedAt"`
	ModifiedBy      string   `json:"modifiedBy"`
	Id              string   `json:"id"`
	FieldNames      []string `json:"fieldNames"`
}
