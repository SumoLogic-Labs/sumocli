package reset

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

func NewCmdUserResetPassword() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "reset password",
		Short: "Initiates a password reset for a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User reset password request started.")
			userResetPassword(id, logger)
			logger.Debug().Msg("User reset password request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user which requires a password reset.")
	cmd.MarkFlagRequired("id")

	return cmd
}

func userResetPassword(id string, logger zerolog.Logger) {
	requestUrl := "v1/users/" + id + "/password/reset"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	logging.LogError(err, logger)

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		jsonErr := json.Unmarshal(responseBody, &responseError)
		fmt.Println(responseBody)
		logging.LogError(jsonErr, logger)
		if responseError.Errors[0].Message != "" {
			fmt.Println(responseError.Errors[0].Message)
		}
	} else {
		fmt.Println("Password reset request completed.")
	}
}
