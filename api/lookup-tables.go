package api

type LookupTableResponse struct {
	CreatedAt       string              `json:"createdAt"`
	CreatedBy       string              `json:"createdBy"`
	ModifiedAt      string              `json:"modifiedAt"`
	ModifiedBy      string              `json:"modifiedBy"`
	Description     string              `json:"description"`
	Fields          []lookupTableFields `json:"fields"`
	PrimaryKeys     []string            `json:"primaryKeys"`
	Ttl             int                 `json:"ttl"`
	SizeLimitAction string              `json:"sizeLimitAction"`
	Name            string              `json:"name"`
	ParentFolderId  string              `json:"parentFolderId"`
	Id              string              `json:"id"`
	ContentPath     string              `json:"contentPath"`
	Size            int                 `json:"size"`
}

type lookupTableFields struct {
	FieldName string `json:"fieldName"`
	FieldType string `json:"fieldType"`
}
