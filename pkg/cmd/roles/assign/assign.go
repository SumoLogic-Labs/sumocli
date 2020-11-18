package assign

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

func NewCmdRoleAssign() *cobra.Command {
	var (
		roleId string
		userId string
		output string
	)

	cmd := &cobra.Command{
		Use:   "assign",
		Short: "Assigns the specified Sumo Logic user to the role.",
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
			logger.Debug().Msg("Role assign request started.")
			assignRole(roleId, userId, output, logger)
			logger.Debug().Msg("Role assign request finished.")
		},
	}

	cmd.Flags().StringVar(&roleId, "roleid", "", "Specify the id of the role")
	cmd.Flags().StringVar(&userId, "userid", "", "Specify the id of the user to remove")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func assignRole(roleId string, userId string, output string, logger zerolog.Logger) {
	var roleInfo api.RoleData

	if roleId == "" || userId == "" {
		fmt.Println("--roleid and --userid fields need to be set")
		os.Exit(0)
	}

	requestUrl := "v1/roles/" + roleId + "/users/" + userId
	client, request := factory.NewHttpRequest("PUT", requestUrl)
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
