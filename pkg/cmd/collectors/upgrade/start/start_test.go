package start

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestUpgradeStart(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://api.au.sumologic.com/api/v1/collectors/upgrades",
		httpmock.NewStringResponder(202, `
{
    "id": "12345",
    "link": {
        "rel": "self",
        "href": "/v1/collectors/upgrades/12345"
    }
}
`))
	upgradeStart(111222, "19.152-1")
}
