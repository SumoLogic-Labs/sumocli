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

type LookupTableStatusResponse struct {
	JobId             string                      `json:"jobId"`
	Status            string                      `json:"status"`
	StatusMessages    []string                    `json:"statusMessages"`
	Errors            []lookupTableStatusErrors   `json:"errors"`
	Warnings          []lookupTableStatusWarnings `json:"warnings"`
	LookupContentId   string                      `json:"lookupContentId"`
	LookupName        string                      `json:"lookupName"`
	LookupContentPath string                      `json:"lookupContentPath"`
	RequestType       string                      `json:"requestType"`
	UserId            string                      `json:"userId"`
	CreatedAt         string                      `json:"createdAt"`
	ModifiedAt        string                      `json:"modifiedAt"`
}

type lookupTableStatusErrors struct {
	Code    string                `json:"code"`
	Message string                `json:"message"`
	Detail  string                `json:"detail"`
	Meta    lookupTableStatusMeta `json:"meta"`
}

type lookupTableStatusMeta struct {
	MinLength    int `json:"minLength"`
	ActualLength int `json:"actualLength"`
}

type lookupTableStatusWarnings struct {
	Message string `json:"message"`
	Cause   string `json:"cause"`
}
