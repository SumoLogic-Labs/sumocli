package api

type DashboardSyncDefinition struct {
	Type        string                      `json:"type"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	DetailLevel int                         `json:"detailLevel"`
	Properties  string                      `json:"properties"`
	Panels      []reportPanelSyncDefinition `json:"panels"`
	Filters     filtersSyncDefinition       `json:"filters"`
}

type FolderSyncDefinition struct {
	Type        string                  `json:"type"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Children    []contentSyncDefinition `json:"children"`
}

type GetContentResponse struct {
	CreatedAt   string   `json:"createdAt"`
	CreatedBy   string   `json:"createdBy"`
	ModifiedAt  string   `json:"modifiedAt"`
	ModifiedBy  string   `json:"modifiedBy"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	ItemType    string   `json:"itemType"`
	ParentId    string   `json:"parentId"`
	Permissions []string `json:"permissions"`
}

type GetPathResponse struct {
	Path string `json:"path"`
}

type MetricsSavedSearchSyncDefinition struct {
	Type
}

type MetricsSearchSyncDefinition struct {
	Type
}

type LookupTableSyncDefinition struct {
	Type
}

type MewboardSyncDefinition struct {
	Type             string                  `json:"type"`
	Name             string                  `json:"name"`
	Description      string                  `json:"description"`
	Title            string                  `json:"title"`
	RootPanel        rootPanelDefinition     `json:"rootPanel"`
	Theme            string                  `json:"theme"`
	TopologyLabelMap topologyLabelMap        `json:"topologyLabelMap"`
	RefreshInterval  int                     `json:"refreshInterval"`
	TimeRange        timeRangeDefinition     `json:"timeRange"`
	Layout           layout                  `json:"layout"`
	Panels           panelsDefinition        `json:"panels"`
	Variables        variablesDefinition     `json:"variables"`
	ColoringRules    coloringRulesDefinition `json:"coloringRules"`
}

type SavedSearchWithScheduleSyncDefinition struct {
	Type
}

type StartExportResponse struct {
	Id string `json:"id"`
}

type ExportStatusResponse struct {
	Status        string      `json:"status"`
	StatusMessage string      `json:"statusMessage,omitempty"`
	Error         exportError `json:"error,omitempty"`
}

type autoComplete struct {
	AutoCompleteType   string               `json:"autoCompleteType"`
	AutoCompleteKey    string               `json:"autoCompleteKey"`
	AutoCompleteValues []autoCompleteValues `json:"autoCompleteValues"`
	LookupFileName     string               `json:"lookupFileName"`
	LookupLabelColumn  string               `json:"lookupLabelColumn"`
	LookupValueColumn  string               `json:"lookupValueColumn"`
}

type autoCompleteValues struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type autoParsingInfo struct {
	Mode string `json:"mode"`
}

type coloringRulesDefinition struct {
	Scope                           string          `json:"scope"`
	SingleSeriesAggregateFunction   string          `json:"singleSeriesAggregateFunction"`
	MultipleSeriesAggregateFunction string          `json:"multipleSeriesAggregateFunction"`
	ColorThresholds                 colorThresholds `json:"colorThresholds"`
}

type colorThresholds struct {
	Color string `json:"color"`
	Min   int    `json:"min"`
	Max   int    `json:"max"`
}

type contentSyncDefinition struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type exportError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

type filtersSyncDefinition struct {
	FieldName    string   `json:"fieldName"`
	Label        string   `json:"label"`
	DefaultValue string   `json:"defaultValue"`
	FilterType   string   `json:"filterType"`
	Properties   string   `json:"properties"`
	PanelIds     []string `json:"panelIds"`
}

type layout struct {
	LayoutType       string            `json:"layouType"`
	LayoutStructures []layoutStructure `json:"layoutStructures"`
}

type layoutStructure struct {
	Key       string `json:"key"`
	Structure string `json:"structure"`
}

type metricsQueriesDefinition struct {
	Query string `json:"query"`
	RowId string `json:"rowId"`
}

type panelsDefinition struct {
	Id                                     string `json:"id"`
	Key                                    string `json:"key"`
	Title                                  string `json:"title"`
	visualSettings                         string `json:"visualSettings"`
	KeepVisualSettingsConsistentWithParent bool   `json:"keepVisualSettingsConsistentWithParent"`
	PanelType                              string `json:"panelType"`
}

type rootPanelDefinition struct {
	Id                                     string                    `json:"id"`
	Key                                    string                    `json:"key"`
	Title                                  string                    `json:"title"`
	VisualSettings                         string                    `json:"visualSettings"`
	KeepVisualSettingsConsistentWithParent bool                      `json:"keepVisualSettingsConsistentWithParent"`
	PanelType                              string                    `json:"panelType"`
	Layout                                 layout                    `json:"layout"`
	Panels                                 []panelsDefinition        `json:"panels"`
	Variables                              []variablesDefinition     `json:"variables"`
	ColoringRules                          []coloringRulesDefinition `json:"coloringRules"`
}

type reportPanelSyncDefinition struct {
	Name                      string                     `json:"name"`
	ViewerType                string                     `json:"viewerType"`
	DetailLevel               int                        `json:"detailLevel"`
	QueryString               string                     `json:"queryString"`
	MetricsQueries            []metricsQueriesDefinition `json:"metricsQueries"`
	TimeRange                 timeRangeDefinition        `json:"timeRange"`
	X                         int                        `json:"x"`
	Y                         int                        `json:"Y"`
	Width                     int                        `json:"width"`
	Height                    int                        `json:"height"`
	Properties                string                     `json:"properties"`
	Id                        string                     `json:"id"`
	DesiredQuantizationInSecs int                        `json:"desiredQuantizationInSecs"`
	QueryParameters           []queryParameters          `json:"queryParameters"`
	AutoParsingInfo           autoParsingInfo            `json:"autoParsingInfo"`
}

type timeRangeDefinition struct {
	Type string                  `json:"type"`
	From timeRangeFromDefinition `json:"from"`
}

type timeRangeFromDefinition struct {
	Type         string `json:"type"`
	RelativeTime string `json:"relativeTime"`
}

type topologyLabelMap struct {
	Service []string `json:"service"`
}

type queryParameters struct {
	Name         string       `json:"name"`
	Label        string       `json:"label"`
	Description  string       `json:"description"`
	DataType     string       `json:"dataType"`
	Value        string       `json:"value"`
	AutoComplete autoComplete `json:"autoComplete"`
}

type variablesDefinition struct {
	Id               string                    `json:"id"`
	Name             string                    `json:"name"`
	DisplayName      string                    `json:"displayName"`
	DefaultValue     string                    `json:"defaultValue"`
	SourceDefinition variablesSourceDefinition `json:"sourceDefinition"`
	AllowMultiSelect bool                      `json:"allowMultiSelect"`
	IncludeAllOption bool                      `json:"includeAllOption"`
	HideFromUI       bool                      `json:"hideFromUI"`
}

type variablesSourceDefinition struct {
	VariableSourceType string `json:"variableSourceType"`
}
