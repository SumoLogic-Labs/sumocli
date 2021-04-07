package assign

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdRoleAssign() *cobra.Command {
	var (
		roleId string
		userId string
	)

	cmd := &cobra.Command{
		Use:   "assign",
		Short: "Assigns the specified Sumo Logic user to the role.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role assign request started.")
			assignRole(roleId, userId, logger)
			logger.Debug().Msg("Role assign request finished.")
		},
	}

	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the id of the role")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the id of the user to remove")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func assignRole(roleId string, userId string, logger zerolog.Logger) {
	var roleInfo api.RoleData

	requestUrl := "v1/roles/" + roleId + "/users/" + userId
	client, request := factory.NewHttpRequest("PUT", requestUrl)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	logging.LogError(err, logger)

	jsonErr := json.Unmarshal(responseBody, &roleInfo)
	logging.LogError(jsonErr, logger)

	roleInfoJson, err := json.MarshalIndent(roleInfo, "", "    ")
	logging.LogError(err, logger)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		fmt.Println(string(roleInfoJson))
	}
}
