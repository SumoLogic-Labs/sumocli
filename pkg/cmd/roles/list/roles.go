package list

import (
	"encoding/json"
	"fmt"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
	"net/http"
	"net/url"
)

type role struct {
	Data []roleData `json:"data"`
}

type roleData struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	FilterPredicate      string   `json:"filterPredicate"`
	Users                []string `json:"users"`
	Capabilities         []string `json:"capabilities"`
	AutofillDependencies bool     `json:"autofillDependencies"`
	CreatedAt            string   `json:"createdAt"`
	CreatedBy            string   `json:"createdBy"`
	ModifiedAt           string   `json:"modifiedAt"`
	ModifiedBy           string   `json:"modifiedBy"`
	Id                   string   `json:"id"`
	SystemDefined        bool     `json:"systemDefined"`
}

func GetRoleId() {

}

func ListRoleIds(numberOfResults string, name string, output bool) {
	var roleInfo role
	client := util2.GetHttpClient()

	request, err := http.NewRequest("GET", util2.GetApiEndpoint()+"v1/roles", nil)
	request.Header.Add("Authorization", util2.GetApiCredentials())
	request.Header.Add("Content-Type", "application/json")
	util2.LogError(err)

	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if name != "" {
		query.Add("name", name)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	util2.LogError(err)

	jsonErr := json.Unmarshal(responseBody, &roleInfo)
	util2.LogError(jsonErr)

	roleInfoJson, err := json.MarshalIndent(roleInfo.Data, "", "    ")
	util2.LogError(err)

	// Determines if the response should be written to a file or to console
	if output == true {
		util2.OutputToFile(roleInfoJson)
	} else {
		fmt.Println(string(roleInfoJson))
	}
}