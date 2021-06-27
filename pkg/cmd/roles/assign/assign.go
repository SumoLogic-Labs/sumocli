package assign

import (
	"encoding/json"
	"fmt"
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
			assignRole(roleId, userId)
		},
	}

	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the id of the role")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the id of the user to remove")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func assignRole(roleId string, userId string) {
	var roleInfo api.RoleData
	log := logging.GetConsoleLogger()
	requestUrl := "v1/roles/" + roleId + "/users/" + userId
	client, request := factory.NewHttpRequest("PUT", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &roleInfo)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	roleInfoJson, err := json.MarshalIndent(roleInfo, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(roleInfoJson))
	}
}
