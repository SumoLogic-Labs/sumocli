package api

type SumoApiEndpoint struct {
	Name     string
	Code     string
	Endpoint string
}

type SumoAuth struct {
	AccessId  string `json:"accessid"`
	AccessKey string `json:"accesskey"`
	Region    string `json:"region"`
	Endpoint  string `json:"endpoint"`
}
