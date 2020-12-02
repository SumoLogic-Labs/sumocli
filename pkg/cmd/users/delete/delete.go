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
)

func NewCmdUserDelete() *cobra.Command {
	var (
		id         string
		transferTo string
	)

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic user",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User delete request started.")
			deleteUser(id, transferTo, logger)
			logger.Debug().Msg("User delete request finished.")
		},
	}

	cmd.PersistentFlags().StringVar(&id, "id", "", "Specify the id of the user to delete.")
	cmd.Flags().StringVar(&transferTo, "transferto", "", "Specify the id of the user to transfer data to.")

	return cmd
}

func deleteUser(id string, transferTo string, logger zerolog.Logger) {
	requestUrl := "v1/users/" + id
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
	} else {
		fmt.Println("User was deleted.")
	}
}
