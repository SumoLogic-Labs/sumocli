package list

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestListCollectors(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.au.sumologic.com/api/v1/collectors",
		httpmock.NewStringResponder(200, `
{  
   "collectors":[  
      {  
         "id":12345,
         "name":"OtherCollector",
         "collectorType":"Installable",
         "alive":true,
         "links":[  
            {  
               "rel":"sources",
               "href":"/v1/collectors/12345/sources"
            }
         ],
         "collectorVersion":"19.33-28",
         "ephemeral":false,
         "description":"Local Windows Collection",
         "osName":"Windows 7",
         "osArch":"amd64",
         "osVersion":"6.1",
         "category":"local"
      }
   ]
}
`))
	listCollectors("hosted", 1000, "", false, true)
}

func TestListCollectorsOffline(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.au.sumologic.com/api/v1/collectors/offline",
		httpmock.NewStringResponder(200, `
{
  "collectors":[{
    "id":12345,
    "name":"My Installed Collector",
    "timeZone":"Etc/UTC",
    "fields":{
      "_budget":"test_budget"
    },
    "links":[{
      "rel":"sources",
      "href":"/v1/collectors/12345/sources"
    }],
    "ephemeral":false,
    "targetCpu":-1,
    "sourceSyncMode":"UI",
    "collectorType":"Installable",
    "collectorVersion":"19.162-12",
    "osVersion":"10.12.6",
    "osName":"Mac OS X",
    "osArch":"x86_64",
    "lastSeenAlive":1521143016128,
    "alive":false
  }]
}
`))
	listCollectors("hosted", 1000, "", true, true)
}
