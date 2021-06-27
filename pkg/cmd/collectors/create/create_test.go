package create

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestCollector(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://api.au.sumologic.com/api/v1/collectors",
		httpmock.NewStringResponder(200, `
{
  "collector":{
    "id":12345,
    "name":"My Hosted Collector",
    "description":"An example Hosted Collector",
    "category":"HTTP Collection",
    "timeZone":"UTC",
    "fields":{
      "_budget":"test_budget"
    },
    "links":[{
      "rel":"sources",
      "href":"/v1/collectors/12345/sources"
    }],
    "collectorType":"Hosted",
    "collectorVersion":"",
    "lastSeenAlive":1536618284387,
    "alive":true
  }
}
`))
	collector := Collector("My Hosted Collector", "An example Hosted Collector", "HTTP Collection", "")
	fmt.Println(collector)
}
