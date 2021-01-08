package api

type CreateSourcesRequest struct {
	SourceType string   `json:"sourceType"`
	Name       string   `json:"name"`
	Interval   int      `json:"interval"`
	HostName   string   `json:"hostName"`
	Metrics    []string `json:"metrics,omitempty"`
}

type SourcesResponse struct {
	Id                         int             `json:"id"`
	Name                       string          `json:"name"`
	HostName                   string          `json:"hostName"`
	AutomaticDateParsing       bool            `json:"automaticDateParsing"`
	MultilineProcessingEnabled bool            `json:"multilineProcessingEnabled"`
	UseAutolineMatching        bool            `json:"useAutolineMatching"`
	ContentType                string          `json:"contentType"`
	ForceTimezone              bool            `json:"forceTimezome"`
	Filters                    []sourceFilters `json:"filters"`
	CutoffTimestamp            int             `json:"cutoffTimestamp"`
	Encoding                   string          `json:"encoding"`
	Interval                   int             `json:"interval"`
	Metrics                    []string        `json:"metrics"`
	SourceType                 string          `json:"sourceType"`
	Alive                      bool            `json:"alive"`
}

type sourceFilters struct {
	FilterType string `json:"filterType"`
	Name       string `json:"name"`
	Regexp     string `json:"regexp"`
}
