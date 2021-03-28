package get

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestGetCollectorUpgrades(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.au.sumologic.com/api/v1/collectors/upgrades/collectors",
		httpmock.NewStringResponder(200, `
{  
   "collectors":[  
      {  
         "id":11111,
         "name":"testcollector",
         "collectorType":"Installable",
         "alive":true,
         "links":[  
            {  
               "rel":"sources",
               "href":"/v1/collectors/1/sources"
            }
         ],
         "collectorVersion":"19.11.11",
         "ephemeral":false,
         "description":"Test Collection",
         "osName":"OS",
         "osArch":"amd64",
         "osVersion":"6.1",
         "category":"local"
      }
   ]
}
`))
	getUpgradableCollectors("", 0, 50)
}
