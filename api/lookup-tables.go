package api

type CreateLookupTableRequest struct {
	Description     string              `json:"description"`
	Fields          []LookupTableFields `json:"fields"`
	PrimaryKeys     []string            `json:"primaryKeys"`
	Ttl             int                 `json:"ttl"`
	SizeLimitAction string              `json:"sizeLimitAction"`
	Name            string              `json:"name"`
	ParentFolderId  string              `json:"parentFolderId"`
}

type EditLookupTable struct {
	Description     string `json:"description"`
	Ttl             int    `json:"ttl"`
	SizeLimitAction string `json:"sizeLimitAction"`
}

type LookupTableResponse struct {
	CreatedAt       string              `json:"createdAt"`
	CreatedBy       string              `json:"createdBy"`
	ModifiedAt      string              `json:"modifiedAt"`
	ModifiedBy      string              `json:"modifiedBy"`
	Description     string              `json:"description"`
	Fields          []LookupTableFields `json:"fields"`
	PrimaryKeys     []string            `json:"primaryKeys"`
	Ttl             int                 `json:"ttl"`
	SizeLimitAction string              `json:"sizeLimitAction"`
	Name            string              `json:"name"`
	ParentFolderId  string              `json:"parentFolderId"`
	Id              string              `json:"id"`
	ContentPath     string              `json:"contentPath"`
	Size            int                 `json:"size"`
}

type LookupTableFields struct {
	FieldName string `json:"fieldName"`
	FieldType string `json:"fieldType"`
}

type LookupTableRequestId struct {
	Id string `json:"id"`
}
