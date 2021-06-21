package api

type CreateMonitorFolder struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type GetUsageInfo struct {
	MonitorType string `json:"monitorType"`
	Usage       int    `json:"usage"`
	Limit       int    `json:"limit"`
	Total       int    `json:"total"`
}

type RootMonitorFolder struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Version     int             `json:"version"`
	CreatedAt   string          `json:"createdAt"`
	CreatedBy   string          `json:"createdBy"`
	ModifiedAt  string          `json:"modifiedAt"`
	ModifiedBy  string          `json:"modifiedBy"`
	ParentId    string          `json:"parentId"`
	ContentType string          `json:"contentType"`
	Type        string          `json:"type"`
	IsSystem    bool            `json:"isSystem"`
	IsMutable   bool            `json:"isMutable"`
	Permissions []string        `json:"permissions"`
	Children    []MonitorFolder `json:"children"`
}

type MonitorFolder struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Version     int      `json:"version"`
	CreatedAt   string   `json:"createdAt"`
	CreatedBy   string   `json:"createdBy"`
	ModifiedAt  string   `json:"modifiedAt"`
	ModifiedBy  string   `json:"modifiedBy"`
	ParentId    string   `json:"parentId"`
	ContentType string   `json:"contentType"`
	Type        string   `json:"type"`
	IsSystem    bool     `json:"isSystem"`
	IsMutable   bool     `json:"isMutable"`
	Permissions []string `json:"permissions"`
}
