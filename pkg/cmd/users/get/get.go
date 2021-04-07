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

func NewCmdGetUser() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic user",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User get request started.")
			getUser(id, logger)
			logger.Debug().Msg("User get request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user to get")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getUser(id string, logger zerolog.Logger) {
	var userInfo api.UserResponse

	requestUrl := "v1/users/" + id
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	logging.LogError(err, logger)

	jsonErr := json.Unmarshal(responseBody, &userInfo)
	logging.LogError(jsonErr, logger)

	userInfoJson, err := json.MarshalIndent(userInfo, "", "    ")
	logging.LogError(err, logger)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		fmt.Println(string(userInfoJson))
	}
}
