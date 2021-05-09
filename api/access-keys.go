package api

type ListAccessKeysResponse struct {
	Data []GetAccessKeysResponse `json:"data"`
}

type GetAccessKeysResponse struct {
	Id          string   `json:"id"`
	Label       string   `json:"label"`
	CorsHeaders []string `json:"corsHeaders"`
	Disabled    bool     `json:"disabled"`
	CreatedAt   string   `json:"createdAt"`
	CreatedBy   string   `json:"createdBy"`
	ModifiedAt  string   `json:"modifiedAt"`
}
