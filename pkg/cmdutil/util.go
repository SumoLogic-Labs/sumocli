package cmdutil

import (
	"net/http"
	"time"
)

func GetHttpClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return client
}
