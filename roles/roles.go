package roles

import (
	"fmt"
	"github.com/wizedkyle/sumocli/util"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetRoleId() {

}

func ListRoleIds(numberOfResults string, name string) {
	client := util.GetHttpClient()

	// TODO: Look at making the creation of the request a function that can be used across the whole code base
	request, err := http.NewRequest("GET", util.GetApiEndpoint()+"v1/roles", nil)
	request.Header.Add("Authorization", util.GetApiCredentials())
	request.Header.Add("Content-Type", "application/json")
	util.LogError(err)

	// if the vars have values create the query url
	query := url.Values{}
	query.Add("limit", numberOfResults)
	query.Add("name", name)
	request.URL.RawQuery = query.Encode()
	fmt.Println(request.URL)

	response, err := client.Do(request)
	util.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	responseString := string(responseBody)
	fmt.Println(responseString)

}
