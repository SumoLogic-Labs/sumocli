package remove

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io/ioutil"
)

func NewCmdRoleRemoveUser() *cobra.Command {
	var (
		roleId string
		userId string
	)

	cmd := &cobra.Command{
		Use:   "remove user",
		Short: "Removes the specified Sumo Logic user from the role.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role remove request started.")
			removeUserRole(roleId, userId, logger)
			logger.Debug().Msg("Role remove request finished.")
		},
	}

	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the id of the role")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the id of the user to remove")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func removeUserRole(roleId string, userId string, logger zerolog.Logger) {
	requestUrl := "v1/roles/" + roleId + "/users/" + userId
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	logging.LogError(err, logger)

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		jsonErr := json.Unmarshal(responseBody, &responseError)
		logging.LogError(jsonErr, logger)
		fmt.Println(responseError.Errors[0].Message)
	} else {
		fmt.Println("User: " + userId + " was removed from role: " + roleId)
	}
}
