package builds

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestGetCollectorUpgrades(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.au.sumologic.com/api/v1/collectors/upgrades/targets",
		httpmock.NewStringResponder(200, ``))
	getBuilds()
}
