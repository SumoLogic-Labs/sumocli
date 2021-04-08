package api

type GetContentResponse struct {
	CreatedAt   string   `json:"createdAt"`
	CreatedBy   string   `json:"createdBy"`
	ModifiedAt  string   `json:"modifiedAt"`
	ModifiedBy  string   `json:"modifiedBy"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	ItemType    string   `json:"itemType"`
	ParentId    string   `json:"parentId"`
	Permissions []string `json:"permissions"`
}

type GetPathResponse struct {
	Path string `json:"path"`
}

type StartExportResponse struct {
	Id string `json:"id"`
}

type ExportStatusResponse struct {
	Status        string      `json:"status"`
	StatusMessage string      `json:"statusMessage,omitempty"`
	Error         exportError `json:"error,omitempty"`
}

type exportError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}
