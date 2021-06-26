package delete

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdRoleDelete() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			deleteRole(id)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to delete")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteRole(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/roles/" + id
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
		if responseError.Errors[0].Code == "acl:role_has_users" {
			fmt.Println("The role wasn't deleted as users are assigned to it." +
				" Please run sumocli roles remove user then re-run sumocli roles delete")
		}
	} else {
		fmt.Println("Role was deleted.")
	}
}
