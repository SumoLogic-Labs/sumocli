package roles

import (
	"encoding/json"
	"fmt"
	"github.com/wizedkyle/sumocli/util"
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

func ListRoleIds(numberOfResults string, name string) {
	var roleInfo role
	client := util.GetHttpClient()

	// TODO: Look at making the creation of the request a function that can be used across the whole code base
	request, err := http.NewRequest("GET", util.GetApiEndpoint()+"v1/roles", nil)
	request.Header.Add("Authorization", util.GetApiCredentials())
	request.Header.Add("Content-Type", "application/json")
	util.LogError(err)

	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if name != "" {
		query.Add("name", name)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	util.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	util.LogError(err)

	jsonErr := json.Unmarshal(responseBody, &roleInfo)
	util.LogError(jsonErr)

	// TODO: Add a flag to output to file or console, there needs to be a conditional IF statement
	for i, _ := range roleInfo.Data {
		roleInfoJson, err := json.MarshalIndent(roleInfo.Data[i], "", "    ")
		util.LogError(err)
		fmt.Print(string(roleInfoJson))
		i++
	}
}
