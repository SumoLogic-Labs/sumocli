package unlock

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

func NewCmdUnlockUser() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "unlock",
		Short: "Unlocks a Sumo Logic user account",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User unlock request started.")
			unlockUser(id, logger)
			logger.Debug().Msg("User unlock request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user account to unlock")

	return cmd
}

func unlockUser(id string, logger zerolog.Logger) {
	if id == "" {
		fmt.Println("--id field needs to be specified.")
		os.Exit(0)
	}

	requestUrl := "v1/users/" + id + "/unlock"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	logging.LogError(err, logger)

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		jsonErr := json.Unmarshal(responseBody, &responseError)
		logging.LogError(jsonErr, logger)
	} else {
		fmt.Println("User account was unlocked.")
	}
}
