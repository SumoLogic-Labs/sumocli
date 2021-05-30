package api

type CreatePartition struct {
	Name              string `json:"name"`
	RoutingExpression string `json:"routingExpression"`
	AnalyticsTier     string `json:"analyticsTier"`
	RetentionPeriod   int    `json:"retentionPeriod"`
	IsCompliant       bool   `json:"isCompliant"`
}

type GetPartitions struct {
	Data []Partitions `json:"data"`
}

type Partitions struct {
	Name                 string `json:"name"`
	RoutingExpression    string `json:"routingExpression"`
	AnalyticsTier        string `json:"analyticsTier"`
	RetentionPeriod      int    `json:"retentionPeriod"`
	IsCompliant          bool   `json:"isCompliant"`
	CreatedAt            string `json:"createdAt"`
	CreatedBy            string `json:"createdBy"`
	ModifiedAt           string `json:"modifiedAt"`
	ModifiedBy           string `json:"modifiedBy"`
	Id                   string `json:"id"`
	TotalBytes           int    `json:"totalBytes"`
	IsActive             bool   `json:"isActive"`
	NewRetentionPeriod   int    `json:"newRetentionPeriod"`
	IndexType            string `json:"indexType"`
	RetentionEffectiveAt string `json:"retentionEffectiveAt"`
	DataForwardingId     string `json:"dataForwardingId"`
}

type UpdatePartition struct {
	RetentionPeriod                  int    `json:"retentionPeriod"`
	ReduceRetentionPeriodImmediately bool   `json:"reduceRetentionPeriodImmediately"`
	IsCompliant                      bool   `json:"isCompliant"`
	RoutingExpression                string `json:"routingExpression"`
}
