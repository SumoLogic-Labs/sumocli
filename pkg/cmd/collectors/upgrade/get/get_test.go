package get

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestGetCollectorUpgrades(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.au.sumologic.com/api/v1/collectors/upgrades/collectors",
		httpmock.NewStringResponder(200, ``))
	getUpgradableCollectors("", 0, 50)
}
