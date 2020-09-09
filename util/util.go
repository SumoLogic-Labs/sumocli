package util

import (
	"encoding/base64"
	"net/http"
	"time"
)

// Sets global variables that are used across multiple packages
var SumoApiEndpoints = map[string]string{
	"au":  "https://api.au.sumologic.com/api/",
	"ca":  "https://api.ca.sumologic.com/api/",
	"de":  "https://api.de.sumologic.com/api/",
	"in":  "https://api.in.sumologic.com/api/",
	"jp":  "https://api.jp.sumologic.com/api/",
	"us1": "https://api.sumologic.com/api/",
	"us2": "https://api.us2.sumologic.com/api/",
}

var (
	AccessId    string
	AccessKey   string
	ApiEndpoint string
)

// GetHttpClient: Creates a HTTP util and returns the util
func GetHttpClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return client
}

func GetApiEndpoint() string {
	apiEndpointUrl := SumoApiEndpoints[ApiEndpoint]
	return apiEndpointUrl
}

func GetApiCredentials() string {
	accessCredentials := AccessId + ":" + AccessKey
	accessCredentialsEnc := base64.StdEncoding.EncodeToString([]byte(accessCredentials))
	accessCredentialsComplete := "Basic " + accessCredentialsEnc
	return accessCredentialsComplete
}
