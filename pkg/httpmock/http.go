package httpmock

import (
	"net/http"
	"net/http/httptest"
)

func mockServer() *httptest.Server {
	handler := http.NewServeMux()
}
