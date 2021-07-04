package config

import (
	"github.com/wizedkyle/sumocli/internal/authentication"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"net/http"
	"time"
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
