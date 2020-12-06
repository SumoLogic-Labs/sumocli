package api

type CreateCollectorRequest struct {
	Collector CollectorCreateData `json:"collector"`
}

type CollectorListResponse struct {
Collectors []CollectorData `json:"collectors"`
}

type CollectorResponse struct{
	Collector CollectorData `json:"collector"`
}

type linkData struct {
Rel string `json:"rel"`
Href string `json:"href"`
}

type CollectorData struct {
Id int64 `json:"id"`
Name string `json:"name"`
CollectorType string `json:"collectorType"`
Alive bool `json:"alive"`
Links []linkData `json:"links,omitempty"`
CollectorVersion string `json:"collectorVersion"`
Ephemeral bool `json:"ephemeral"`
Description string `json:"description,omitempty"`
OsName string `json:"osName,omitempty"`
OsArch string `json:"osArch,omitempty"`
OsVersion string `json:"osVersion,omitempty"`
Category string `json:"category,omitempty"`
CutoffRelativeTime string `json:"cutoffRelativeTime,omitempty"`
CutoffTimestamp int64 `json:"cutoffTimestamp,omitempty"`
HostName string `json:"hostName,omitempty"`
LastSeenAlive int64 `json:"lastSeenAlive,omitempty"`
SourceSyncMode string `json:"sourceSyncMode,omitempty"`
TimeZone string `json:"timeZone,omitempty"`
TargetCpu int64 `json:"targetCpu,omitempty"`
Fields string `json:"fields,omitempty"`
}

type CollectorCreateData struct {
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	CollectorType string `json:"collectorType"`
	Category string `json:"category,omitempty"`
	HostName string `json:"hostName,omitempty"`
	Fields string `json:"fields,omitempty"`
}