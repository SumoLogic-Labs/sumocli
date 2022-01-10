package config

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/internal/build"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"net/http"
	"time"
)

var (
	userAgent = "Sumocli " + build.Version
)

func GetSumoLogicSDKConfig() *cip.APIClient {
	accessId, accessKey, endpoint := authentication.ReadAuthCredentials()
	client := cip.APIClient{
		Cfg: &cip.Configuration{
			Authentication: cip.BasicAuth{
				AccessId:  accessId,
				AccessKey: accessKey,
			},
			BasePath:  endpoint,
			UserAgent: "Sumocli",
			HTTPClient: &http.Client{
				Timeout: time.Second * 20,
			},
		},
	}
	return &client
}

func GetUserAgent() string {
	return userAgent
}
