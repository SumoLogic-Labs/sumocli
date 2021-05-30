package api

type CreateScheduledView struct {
	Query            string `json:"query"`
	IndexName        string `json:"indexName"`
	StartTime        string `json:"startTime"`
	RetentionPeriod  int    `json:"retentionPeriod"`
	DataForwardingId string `json:"dataForwardingId"`
	ParsingMode      string `json:"parsingMode"`
}

type GetScheduledViews struct {
	Data []ScheduledViews `json:"data"`
}

type ScheduledViews struct {
	Query               string `json:"query"`
	IndexName           string `json:"indexName"`
	StartTime           string `json:"startTime"`
	RetentionPeriod     int    `json:"retentionPeriod"`
	DataForwardingId    string `json:"dataForwardingId"`
	ParsingMode         string `json:"parsingMode"`
	Id                  string `json:"id"`
	CreatedAt           string `json:"createdAt"`
	CreatedByOptimizeIt bool   `json:"createdByOptimizeIt"`
	Error               string `json:"error"`
	Status              string `json:"status"`
	TotalBytes          int    `json:"totalBytes"`
	TotalMessageCount   int    `json:"totalMessageCount"`
	CreatedBy           string `json:"createdBy"`
}

type UpdateScheduledView struct {
	DataForwardingId                 string `json:"dataForwardingId"`
	RetentionPeriod                  int    `json:"retentionPeriod"`
	ReduceRetentionPeriodImmediately bool   `json:"reduceRetentionPeriodImmediately"`
}
