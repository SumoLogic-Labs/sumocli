package api

type CreateIngestBudgetRequest struct {
	Name           string `json:"name"`
	FieldValue     string `json:"fieldValue"`
	CapacityBytes  int    `json:"capacityBytes"`
	Timezone       string `json:"timezone"`
	ResetTime      string `json:"resetTime"`
	Description    string `json:"description"`
	Action         string `json:"action"`
	AuditThreshold int    `json:"auditThreshold"`
}

type GetIngestBudget struct {
	Name               string               `json:"name"`
	FieldValue         string               `json:"fieldValue"`
	CapacityBytes      int                  `json:"capacityBytes"`
	TimeZone           string               `json:"timeZone"`
	ResetTime          string               `json:"resetTime"`
	Description        string               `json:"description"`
	Action             string               `json:"action"`
	AuditThreshold     int                  `json:"auditThreshold"`
	CreatedAt          string               `json:"createdAt"`
	CreatedByUser      ingestBudgetUserInfo `json:"createdByUser"`
	ModifiedAt         string               `json:"modifiedAt"`
	ModifiedByUser     ingestBudgetUserInfo `json:"modifiedByUser"`
	Id                 string               `json:"id"`
	UsageBytes         int                  `json:"usageBytes"`
	UsageStatus        string               `json:"usageStatus"`
	NumberOfCollectors int                  `json:"numberOfCollectors"`
}

type ListIngestBudgets struct {
	Data []GetIngestBudget `json:"data"`
}

type ingestBudgetUserInfo struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
