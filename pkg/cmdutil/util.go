package cmdutil

import (
	"net/http"
	"time"
)

func GetHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}
