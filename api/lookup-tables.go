package api

type GetLookupTableResponse struct {
	CreatedAt   string `json:"createdAt"`
	CreatedBy   string `json:"createdBy"`
	ModifiedAt  string `json:"modifiedAt"`
	ModifiedBy  string `json:"modifiedBy"`
	Description string `json:"description"`
}
