package get

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestGetCollectorById(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.au.sumologic.com/api/v1/collectors/12345",
		httpmock.NewStringResponder(200, `
{
  "collector":{
    "id":12345,
    "name":"test",
    "timeZone":"UTC",
    "ephemeral":false,
    "sourceSyncMode":"Json",
    "collectorType":"Installable",
    "collectorVersion":"19.162-12",
    "osVersion":"3.13.0-92-generic",
    "osName":"Linux",
    "osArch":"amd64",
    "lastSeenAlive":1471553524302,
    "alive":true
  }
}
`))
	getCollector(12345, "")
}

func TestGetCollectorByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.au.sumologic.com/api/v1/collectors/name/test",
		httpmock.NewStringResponder(200, `
{
  "collector":{
    "id":12345,
    "name":"test",
    "timeZone":"UTC",
    "ephemeral":false,
    "sourceSyncMode":"Json",
    "collectorType":"Installable",
    "collectorVersion":"19.162-12",
    "osVersion":"3.13.0-92-generic",
    "osName":"Linux",
    "osArch":"amd64",
    "lastSeenAlive":1471553524302,
    "alive":true
  }
}
`))
	getCollector(0, "test")
}
