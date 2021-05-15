package api

type CreateTokenRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Type        string `json:"type"`
}

type GetTokenResponse struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	Version            int    `json:"version"`
	CreatedAt          string `json:"createdAt"`
	CreatedBy          string `json:"createdBy"`
	ModifiedAt         string `json:"modifiedAt"`
	ModifiedBy         string `json:"modifiedBy"`
	EncodedTokenAndUrl string `json:"encodedTokenAndUrl"`
}

type ListTokenResponse struct {
	Data []GetTokenResponse `json:"data"`
}

type UpdateTokenRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Version     int    `json:"version"`
}
