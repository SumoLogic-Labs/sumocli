package list

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io/ioutil"
	"net/url"
	"strings"
)

func NewCmdRoleList() *cobra.Command {
	var (
		numberOfResults string
		filter          string
		output          string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic roles",
		Long: `The following fields can be exported using the --output command:
name
description
filterPredicate
users
capabilities
id
`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role list request started.")
			listRoles(numberOfResults, filter, output, logger)
			logger.Debug().Msg("Role list request finished.")
		},
	}

	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func listRoles(numberOfResults string, name string, output string, logger zerolog.Logger) {
	var roleInfo api.Role

	client, request := factory.NewHttpRequest("GET", "v1/roles")
	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if filter != "" {
		query.Add("name", filter)
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

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody)
	} else {
		if factory.ValidateRoleOutput(output) == true {
			value := gjson.Get(string(roleInfoJson), "#."+output)
			formattedValue := strings.Trim(value.String(), `"[]"`)
			fmt.Println(formattedValue)
		} else {
			fmt.Println(string(roleInfoJson))
		}
	}
}

func validateOutput(output string) bool {
	switch output {
	case
		"name",
		"description",
		"filterPredicate",
		"users",
		"capabilities",
		"id":
		return true
	}
	return false
}
