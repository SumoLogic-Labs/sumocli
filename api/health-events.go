package api

type CreateHealthEventRequest struct {
	Data []HealthEventRequest `json:"data"`
}

type ListHealthEvent struct {
	Data []HealthEvent `json:"data"`
}

type HealthEvent struct {
	EventId          string                      `json:"eventId"`
	EventName        string                      `json:"eventName"`
	Details          healthEventDetails          `json:"details"`
	ResourceIdentity healthEventResourceIdentity `json:"resourceIdentity"`
	EventTime        string                      `json:"eventTime"`
	Subsystem        string                      `json:"subsystem"`
	severityLevel    string                      `json:"severityLevel"`
}

type HealthEventRequest struct {
	CollectorId            string `json:"collectorId,omitempty"`
	CollectorName          string `json:"collectorName,omitempty"`
	Id                     string `json:"id"`
	IngestBudgetFieldValue string `json:"ingestBudgetFieldValue,omitempty"`
	Name                   string `json:"name"`
	Scope                  string `json:"scope,omitempty"`
	Type                   string `json:"type"`
}

type healthEventDetails struct {
	TrackerId   string `json:"trackerId"`
	Error       string `json:"error"`
	Description string `json:"description"`
}

type healthEventResourceIdentity struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
