package api

type EventHubCollection struct {
	ApiVersion string         `json:"api.version"`
	Source     EventHubSource `json:"source"`
	Config     EventHubConfig `json:"config"`
	SourceType string         `json:"sourceType"`
}

type EventHubConfig struct {
	Name                    string            `json:"name"`
	Description             string            `json:"description"`
	Namespace               string            `json:"namespace"`
	HubName                 string            `json:"hub_name"`
	AccessPolicyName        string            `json:"access_policy_name"`
	AccessPolicyKey         string            `json:"access_policy_key"`
	ConsumerGroup           string            `json:"consumer_group"`
	Fields                  map[string]string `json:"fields"`
	category                string            `json:"category"`
	ReceiveWithLatestOffset bool              `json:"receive_with_latest_offset"`
}

type EventHubSource struct {
	SchemaRef EventHubSourceSchema `json:"schemaRef"`
}

type EventHubSourceSchema struct {
	Type string `json:"type"`
}
