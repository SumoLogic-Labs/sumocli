package api

type CollectorsResponse struct {
	Alive              bool              `json:"alive"`
	Category           string            `json:"category,omitempty"`
	CollectorType      string            `json:"collectorType"`
	CollectorVersion   string            `json:"collectorVersion"`
	CutoffRelativeTime string            `json:"cutoffRelativeTime,omitempty"`
	CutoffTimestamp    int               `json:"cutoffTimestamp,omitempty"`
	Description        string            `json:"description,omitempty"`
	Ephemeral          bool              `json:"ephemeral"`
	Fields             map[string]string `json:"fields,omitempty"`
	HostName           string            `json:"hostName,omitempty"`
	Id                 int               `json:"id"`
	LastSeenAlive      int               `json:"lastSeenAlive,omitempty"`
	Name               string            `json:"name"`
	SourceSyncMode     string            `json:"sourceSyncMode,omitempty"`
	TimeZone           string            `json:"timeZone,omitempty"`
	TargetCpu          int               `json:"targetCpu,omitempty"`
	OsName             string            `json:"osName,omitempty"`
	OsVersion          string            `json:"osVersion,omitempty"`
	OsArch             string            `json:"osArch,omitempty"`
	OsTime             int               `json:"osTime,omitempty"`
}
