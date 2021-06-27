package remove

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
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
			removeUserRole(roleId, userId)
		},
	}

	cmd.Flags().StringVar(&roleId, "roleId", "", "Specify the id of the role")
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the id of the user to remove")
	cmd.MarkFlagRequired("roleId")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func removeUserRole(roleId string, userId string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/roles/" + roleId + "/users/" + userId
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		err = json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		fmt.Println(responseError.Errors[0].Message)
	} else {
		fmt.Println("User: " + userId + " was removed from role: " + roleId)
	}
}
