package api

type GetSubdomain struct {
	CreatedAt  string `json:"createdAt"`
	CreatedBy  string `json:"createdBy"`
	ModifiedAt string `json:"modifiedAt"`
	ModifiedBy string `json:"modifiedBy"`
	Subdomain  string `json:"subdomain"`
	URL        string `json:"url"`
}

type UpdateSubdomainRequest struct {
	Subdomain string `json:"subdomain"`
}
