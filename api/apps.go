package api

type ListApps struct {
	Apps []AppDetails `json:"apps"`
}

type AppDetails struct {
	AppDefinition appDefinition `json:"appDefinition"`
	AppManifest   appManifest   `json:"appManifest"`
}

type InstallAppRequest struct {
	Name                string          `json:"name"`
	Description         string          `json:"description"`
	DestinationFolderId string          `json:"destinationFolderId"`
	DataSourceValues    DataSourceValue `json:"dataSourceValues"`
}

type InstallAppResponse struct {
	Id string `json:"id"`
}

type InstallStatusResponse struct {
	Status        string `json:"status"`
	StatusMessage string `json:"statusMessage"`
}

type appDefinition struct {
	ContentId       string `json:"contentId"`
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	AppVersion      string `json:"appVersion"`
	Preview         bool   `json:"preview"`
	ManifestVersion string `json:"manifestVersion"`
}

type appManifest struct {
	Family                           string       `json:"family"`
	Description                      string       `json:"description"`
	Categories                       []string     `json:"categories"`
	HoverText                        string       `json:"hoverText"`
	IconURL                          string       `json:"iconURL"`
	ScreenshotURLs                   []string     `json:"screenshotURLs"`
	HelpURL                          string       `json:"helpURL"`
	HelpDocIdMap                     helpDocIdMap `json:"helpDocIdMap"`
	CommunityURL                     string       `json:"communityURL"`
	Requirements                     []string     `json:"requirements"`
	AccountTypes                     []string     `json:"accountTypes"`
	RequiresInstallationInstructions bool         `json:"requiresInstallationInstructions"`
	InstallationInstructions         string       `json:"installationInstructions"`
	Parameters                       []parameters `json:"parameters"`
	Author                           string       `json:"author"`
	AuthorWebsite                    string       `json:"author_website"`
}

type DataSourceValue struct {
	LogSrc string `json:"logsrc"`
}

type helpDocIdMap struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

type parameters struct {
	ParameterType  string `json:"parameterType"`
	ParameterId    string `json:"parameterId"`
	DataSourceType string `json:"dataSourceType"`
	Label          string `json:"label"`
	Description    string `json:"description"`
	Example        string `json:"example"`
	Hidden         bool   `json:"hidden"`
}
