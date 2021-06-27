package reset

import (
	"encoding/json"
	"fmt"
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
			userResetPassword(id)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user which requires a password reset.")
	cmd.MarkFlagRequired("id")

	return cmd
}

func userResetPassword(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/users/" + id + "/password/reset"
	client, request := factory.NewHttpRequest("POST", requestUrl)
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
		fmt.Println(responseBody)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		if responseError.Errors[0].Message != "" {
			fmt.Println(responseError.Errors[0].Message)
		}
	} else {
		fmt.Println("Password reset request completed.")
	}
}
