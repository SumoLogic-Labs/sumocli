package api

type GetIngestBudgetV2 struct {
	Name           string `json:"name"`
	Scope          string `json:"scope"`
	CapacityBytes  int    `json:"capacityBytes"`
	Timezone       string `json:"timezone"`
	ResetTime      string `json:"resetTime"`
	Description    string `json:"description"`
	Action         string `json:"action"`
	AuditThreshold int    `json:"auditThreshold"`
	Id             string `json:"id"`
	UsageBytes     int    `json:"usageBytes"`
	UsageStatus    string `json:"usageStatus"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
	ModifiedAt     string `json:"modifiedAt"`
	ModifiedBy     string `json:"modifiedBy"`
	BudgetVersion  int    `json:"budgetVersion"`
}

type ListIngestBudgetsV2 struct {
	Data []GetIngestBudgetV2 `json:"data"`
}
