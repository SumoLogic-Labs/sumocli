package api

type Collectors struct {
	Data []CollectorsResponse `json:"collectors"`
}

type CreateCollectorRequest struct {
	Collector CreateCollector `json:"collector"`
}

type CreateCollector struct {
	CollectorType string            `json:"collectorType"`
	Name          string            `json:"name"`
	Description   string            `json:"description,omitempty"`
	Category      string            `json:"category,omitempty"`
	Fields        map[string]string `json:"fields,omitempty"`
}

type CollectorResponse struct {
	Collector CollectorsResponse `json:"collector"`
}

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
	Links              []collectorLinks  `json:"links,omitempty"`
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

type Targets struct {
	Targets []targetsResponse `json:"targets"`
}

type UpgradeTask struct {
	CollectorId int    `json:"collectorId"`
	ToVersion   string `json:"toVersion"`
}

type UpgradeTaskResponse struct {
	Id   string          `json:"id"`
	Link upgradeTaskLink `json:"link"`
}

type UpgradeTaskStatus struct {
	Upgrade upgradeTaskProgress `json:"upgrade"`
}

type collectorLinks struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type targetsResponse struct {
	Version string `json:"version"`
	Latest  bool   `json:"latest"`
}

type upgradeTaskLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type upgradeTaskProgress struct {
	Id          int    `json:"id"`
	CollectorId int    `json:"collectorId"`
	ToVersion   string `json:"toVersion"`
	RequestTime int    `json:"requestTime"`
	Status      int    `json:"status"`
	Message     string `json:"message"`
}
