package get

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

func NewCmdRoleGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic role information",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role get request started.")
			getRole(id, logger)
			logger.Debug().Msg("Role get request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to get")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getRole(id string, logger zerolog.Logger) {
	var roleInfo api.RoleData

	requestUrl := "v1/roles/" + id
	client, request := factory.NewHttpRequest("GET", requestUrl)
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
