package api

type CreateDashboard struct {
	Title            string                    `json:"title"`
	Description      string                    `json:"description"`
	FolderId         string                    `json:"folderId"`
	TopologyLabelMap topologyLabelMap          `json:"topologyLabelMap"`
	Domain           string                    `json:"domain"`
	RefreshInterval  int                       `json:"refreshInterval"`
	TimeRange        timeRangeDefinition       `json:"timeRange"`
	Panels           []panelsDefinition        `json:"panels"`
	Layout           layout                    `json:"layout"`
	Variables        []variablesDefinition     `json:"variables"`
	Theme            string                    `json:"theme"`
	ColoringRules    []coloringRulesDefinition `json:"coloringRules"`
}

type GetDashboard struct {
	Title            string                    `json:"title"`
	Description      string                    `json:"description"`
	FolderId         string                    `json:"folderId"`
	TopologyLabelMap topologyLabelMap          `json:"topologyLabelMap"`
	Domain           string                    `json:"domain"`
	RefreshInterval  int                       `json:"refreshInterval"`
	TimeRange        timeRangeDefinition       `json:"timeRange"`
	Panels           []panelsDefinition        `json:"panels"`
	Layout           layout                    `json:"layout"`
	Variables        []variablesDefinition     `json:"variables"`
	Theme            string                    `json:"theme"`
	ColoringRules    []coloringRulesDefinition `json:"coloringRules"`
	Id               string                    `json:"id"`
}
