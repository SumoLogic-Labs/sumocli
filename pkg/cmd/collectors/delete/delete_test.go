package delete

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestDeleteCollector(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", "https://api.au.sumologic.com/api/v1/collectors/123456789",
		httpmock.NewStringResponder(200, ``))
	deleteCollector(100, false, "123456789", false)
}

func TestDeleteOfflineCollectors(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", "https://api.au.sumologic.com/api/v1/collectors/offline",
		httpmock.NewStringResponder(200, ``))
	deleteCollector(100, true, "", true)
}
