package status

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestUpgradableCollectorStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.au.sumologic.com/api/v1/collectors/upgrades/12345",
		httpmock.NewStringResponder(200, `
{
    "upgrade": {
        "id": 12345,
        "collectorId": 100000000,
        "toVersion": "19.152-1",
        "requestTime": 1465855411044,
        "status": 2,
        "message": ""
    }
} 
`))
}
