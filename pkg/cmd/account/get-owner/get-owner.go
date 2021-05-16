package get_owner

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	GetUser "github.com/wizedkyle/sumocli/pkg/cmd/users/get"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"strings"
)

func NewCmdAccountGetOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-owner",
		Short: "Returns the user identifier of the account owner.",
		Run: func(cmd *cobra.Command, args []string) {
			getOwner()
		},
	}
	return cmd
}

func getOwner() {
	var userId string
	log := logging.GetConsoleLogger()
	requestUrl := "v1/account/accountOwner"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}
	userId = string(responseBody)
	userId = strings.Trim(userId, "\"")
	GetUser.GetUser(userId)
}
