package httpmock

import "net/http"

type RoundTripFunc func(req *http.Client) *http.Response

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
