package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
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

func NewCmdRoleGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			getRoles(id)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to get")

	return cmd
}

func getRoles(id string) {
	var roleInfo role
	client := util2.GetHttpClient()
	authToken, apiEndpoint := login.ReadCredentials()

	request, err := http.NewRequest("GET", apiEndpoint+"v1/roles", nil)
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	util2.LogError(err)

	query := url.Values{}
	query.Add("id", id)
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

	fmt.Println(string(roleInfoJson))
}
