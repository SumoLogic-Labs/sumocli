package api

type CreateHTTPSource struct {
	ApiVersion string     `json:"api.version"`
	Source     HttpSource `json:"source"`
}

type CreateLocalFileSource struct {
	ApiVersion string          `json:"api.version"`
	Source     LocalFileSource `json:"source"`
}

type GetLocalFileSource struct {
	Source LocalFileResponse `json:"source"`
}

type LocalFileResponse struct {
	Id                         int             `json:"id"`
	Name                       string          `json:"name"`
	Category                   string          `json:"category"`
	AutomaticDateParsing       bool            `json:"automaticDateParsing"`
	MultilineProcessingEnabled bool            `json:"multilineProcessingEnabled"`
	UseAutolineMatching        bool            `json:"useAutolineMatching"`
	ForceTimeZone              bool            `json:"forceTimeZone"`
	Filters                    []SourceFilters `json:"filters"`
	CutoffTimestamp            int             `json:"cutoffTimestamp"`
	Encoding                   string          `json:"encoding"`
	PathExpression             string          `json:"pathExpression"`
	Blacklist                  []string        `json:"blacklist"`
	SourceType                 string          `json:"sourceType"`
	Alive                      bool            `json:"alive"`
}

type HttpSource struct {
	SourceType                 string            `json:"sourceType"`
	Name                       string            `json:"name"`
	Category                   string            `json:"category"`
	Fields                     map[string]string `json:"fields"`
	MessagePerRequest          bool              `json:"messagePerRequest"`
	MultilineProcessingEnabled bool              `json:"multilineProcessingEnabled"`
}

type LocalFileSource struct {
	Name                       string          `json:"name"`
	Category                   string          `json:"category"`
	AutomaticDateParsing       bool            `json:"automaticDateParsing"`
	MultilineProcessingEnabled bool            `json:"multilineProcessingEnabled"`
	UseAutolineMatching        bool            `json:"useAutolineMatching"`
	ForceTimeZone              bool            `json:"forceTimeZone"`
	TimeZone                   string          `json:"timeZone"`
	Filters                    []SourceFilters `json:"filters"`
	CutoffRelativeTime         string          `json:"cutoffRelativeTime"`
	Encoding                   string          `json:"encoding"`
	PathExpression             string          `json:"pathExpression"`
	Blacklist                  []string        `json:"blacklist"`
	SourceType                 string          `json:"sourceType"`
}

type ListSources struct {
	Sources []GetSourcesResponse `json:"sources"`
}

type CreateSourceResponse struct {
	Source SourcesResponse `json:"source"`
}

type GetSourcesResponse struct {
	Id                         string          `json:"id"`
	Name                       string          `json:"name"`
	Category                   string          `json:"category"`
	HostName                   string          `json:"hostName"`
	AutomaticDateParsing       bool            `json:"automaticDateParsing"`
	MultilineProcessingEnabled bool            `json:"multilineProcessingEnabled"`
	UseAutolineMatching        bool            `json:"useAutolineMatching"`
	ForceTimeZone              bool            `json:"forceTimeZone"`
	Filters                    []SourceFilters `json:"filters"`
	CutoffTimestamp            int             `json:"cutoffTimestamp"`
	Encoding                   string          `json:"encoding"`
	PathExpression             string          `json:"pathExpression"`
	Blacklist                  []string        `json:"blacklist"`
	SourceType                 string          `json:"sourceType"`
	Alive                      bool            `json:"alive"`
}

type SourcesResponse struct {
	Id                         int               `json:"id"`
	Name                       string            `json:"name"`
	HostName                   string            `json:"hostName,omitempty"`
	AutomaticDateParsing       bool              `json:"automaticDateParsing"`
	MultilineProcessingEnabled bool              `json:"multilineProcessingEnabled"`
	UseAutolineMatching        bool              `json:"useAutolineMatching"`
	ForceTimezone              bool              `json:"forceTimezome"`
	Filters                    []SourceFilters   `json:"filters"`
	CutoffTimestamp            int               `json:"cutoffTimestamp"`
	Encoding                   string            `json:"encoding"`
	Interval                   int               `json:"interval"`
	Metrics                    []string          `json:"metrics"`
	SourceType                 string            `json:"sourceType"`
	Alive                      bool              `json:"alive"`
	Url                        string            `json:"url,omitempty"`
	Category                   string            `json:"category,omitempty"`
	Fields                     map[string]string `json:"fields,omitempty"`
	MessagePerRequest          bool              `json:"messagePerRequest"`
}

type UpdateSourcesResponse struct {
	Id                         string            `json:"id"`
	Name                       string            `json:"name"`
	HostName                   string            `json:"hostName,omitempty"`
	AutomaticDateParsing       bool              `json:"automaticDateParsing"`
	MultilineProcessingEnabled bool              `json:"multilineProcessingEnabled"`
	UseAutolineMatching        bool              `json:"useAutolineMatching"`
	ForceTimezone              bool              `json:"forceTimezome"`
	Filters                    []SourceFilters   `json:"filters"`
	CutoffTimestamp            int               `json:"cutoffTimestamp"`
	Encoding                   string            `json:"encoding"`
	Interval                   int               `json:"interval"`
	Metrics                    []string          `json:"metrics"`
	SourceType                 string            `json:"sourceType"`
	Alive                      bool              `json:"alive"`
	Url                        string            `json:"url,omitempty"`
	Category                   string            `json:"category,omitempty"`
	Fields                     map[string]string `json:"fields,omitempty"`
	MessagePerRequest          bool              `json:"messagePerRequest"`
}

type SourceFilters struct {
	FilterType string `json:"filterType"`
	Name       string `json:"name"`
	Regexp     string `json:"regexp"`
}
