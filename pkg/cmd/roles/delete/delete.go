package delete

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io/ioutil"
	"os"
)

func NewCmdRoleDelete() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role delete request started.")
			deleteRole(id, logger)
			logger.Debug().Msg("Role delete request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to delete")

	return cmd
}

func deleteRole(id string, logger zerolog.Logger) {
	if id == "" {
		fmt.Println("--id field needs to be set.")
		os.Exit(0)
	}

	requestUrl := "v1/roles/" + id
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
		if responseError.Errors[0].Code == "acl:role_has_users" {
			fmt.Println("The role wasn't deleted as users are assigned to it." +
				" Please run sumocli roles remove user then re-run sumocli roles delete")
		}
	} else {
		fmt.Println("Role was deleted.")
	}
}
