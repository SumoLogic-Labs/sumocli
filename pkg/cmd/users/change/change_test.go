package change

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

func TestUserChangeEmail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder("POST", "",
		func(request *http.Request) (*http.Response, error) {
			response := httpmock.NewStringResponse(200, `
{
}
`,
			)
			response.Header.Add("Content-Type", "application/json")
			return response, nil
		})
}
