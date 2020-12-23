package get

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
	"os"
	"strings"
)

func NewCmdRoleGet() *cobra.Command {
	var (
		id     string
		output string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic role information",
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
			logger.Debug().Msg("Role get request started.")
			getRole(id, output, logger)
			logger.Debug().Msg("Role get request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to get")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func getRole(id string, output string, logger zerolog.Logger) {
	var roleInfo api.RoleData

	if id == "" {
		fmt.Println("--id field needs to be specified.")
		os.Exit(0)
	}

	requestUrl := "v1/roles/" + id
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	logging.LogError(err, logger)

	jsonErr := json.Unmarshal(responseBody, &roleInfo)
	logging.LogError(jsonErr, logger)

	roleInfoJson, err := json.MarshalIndent(roleInfo, "", "    ")
	logging.LogError(err, logger)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		if factory.ValidateRoleOutput(output) == true {
			value := gjson.Get(string(roleInfoJson), output)
			formattedValue := strings.Trim(value.String(), `"[]"`)
			fmt.Println(formattedValue)
		} else {
			fmt.Println(string(roleInfoJson))
		}
	}
}
