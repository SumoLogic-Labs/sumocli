package factory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/api"
	"github.com/SumoLogic-Labs/sumocli/config"
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/logging"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

func newHttpClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return client
}

func NewLiveTailHttpRequest(method string, liveTailEndpoint string, body []byte) (*http.Client, *http.Request) {
	client := newHttpClient()
	request, _ := http.NewRequest(method, liveTailEndpoint, bytes.NewBuffer(body))
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", config.GetUserAgent())
	return client, request
}

func StartLiveTailHttpRequest(method string, liveTailEndpoint string) (*http.Client, *http.Request) {
	client := newHttpClient()
	request, _ := http.NewRequest(method, liveTailEndpoint, nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", config.GetUserAgent())
	return client, request
}

func NewHttpRequest(method string, apiUrl string) (*http.Client, *http.Request) {
	client := newHttpClient()
	authToken, endpoint := authentication.ReadCredentials()
	request, _ := http.NewRequest(method, endpoint+apiUrl, nil)
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", config.GetUserAgent())
	return client, request
}

func NewHttpRequestWithBody(method string, apiUrl string, body []byte) (*http.Client, *http.Request) {
	client := newHttpClient()
	authToken, endpoint := authentication.ReadCredentials()
	request, _ := http.NewRequest(method, endpoint+apiUrl, bytes.NewBuffer(body))
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", config.GetUserAgent())
	return client, request
}

func HttpError(statusCode int, errorMessage []byte, logger zerolog.Logger) {
	if statusCode == 400 {
		var responseError api.ResponseError
		jsonErr := json.Unmarshal(errorMessage, &responseError)
		logging.LogError(jsonErr, logger)
		fmt.Println(responseError.Errors[0].Message)
	} else if statusCode == 401 {
		fmt.Println("Unauthorized access please check the user exists, and credentials are valid.")
	} else if statusCode == 403 {
		fmt.Println("This operation is not allowed for your account type or the user doesn't have the role capability to perform this action.")
	} else if statusCode == 404 {
		fmt.Println("Requested resource could not be found.")
	} else if statusCode == 405 {
		fmt.Println("Unsupported method for URL.")
	} else if statusCode == 415 {
		fmt.Println("Invalid content type.")
	} else if statusCode == 429 {
		fmt.Println("The API request rate is higher than 4 request per second or inflight API requests are higher than 10 requests per second.")
	} else if statusCode == 500 {
		fmt.Println("Internal server error.")
	} else if statusCode == 503 {
		fmt.Println("Service is currently unavailable.")
	}
}
