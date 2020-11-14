package list

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/login"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"github.com/wizedkyle/sumocli/pkg/logging"
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

func NewCmdRoleList() *cobra.Command {
	var (
		numberOfResults string
		filter          string
		output          bool
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic roles",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role list request started.")
			roles(numberOfResults, filter, output, logger)
			logger.Debug().Msg("Role list request finished.")
		},
	}

	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")
	cmd.Flags().BoolVar(&output, "output", false, "Output results to a file, defaults to false")

	return cmd
}

func roles(numberOfResults string, name string, output bool, logger zerolog.Logger, ) {
	var roleInfo role
	client := util2.GetHttpClient()
	authToken, apiEndpoint := login.ReadCredentials()

	request, err := http.NewRequest("GET", apiEndpoint+"v1/roles", nil)
	request.Header.Add("Authorization", authToken)
	request.Header.Add("Content-Type", "application/json")
	logging.LogErrorWithMessage("Creating authorization header failed, please review the credentials supplied in sumocli login.", err, logger)

	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if name != "" {
		query.Add("name", name)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	logging.LogErrorWithMessage("Authorization was not successful, please review your connectivity and credentials.", err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	logging.LogErrorWithMessage("Reading the response body was not successful.", err, logger)

	jsonErr := json.Unmarshal(responseBody, &roleInfo)
	logging.LogErrorWithMessage("Parsing the response body as JSON was not successful.", jsonErr, logger)

	roleInfoJson, err := json.MarshalIndent(roleInfo.Data, "", "    ")
	logging.LogErrorWithMessage("Formatting the role info as JSON was not successful.", jsonErr, logger)

	// Determines if the response should be written to a file or to console
	// TODO we may be able to use zerolog for this
	if output == true {
		util2.OutputToFile(roleInfoJson)
	} else {
		fmt.Println(string(roleInfoJson))
	}
}
