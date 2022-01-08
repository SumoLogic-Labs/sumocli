package remove_allowlist_user

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/api"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Labs/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdSamlRemoveAllowListUser() *cobra.Command {
	var userId string

	cmd := &cobra.Command{
		Use:   "remove-allowlist-user",
		Short: "Remove an allowlisted user requiring them to sign in using SAML.",
		Run: func(cmd *cobra.Command, args []string) {
			removeAllowListUser(userId)
		},
	}
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the id of the user")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func removeAllowListUser(userId string) {
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/saml/allowlistedUsers/" + userId
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		err := json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
	} else {
		fmt.Println("User was successfully removed form the allowlist for SAML lockdown.")
	}
}
