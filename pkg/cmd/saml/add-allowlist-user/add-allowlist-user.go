package add_allowlist_user

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/api"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Incubator/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdSamlAddAllowListUser() *cobra.Command {
	var userId string

	cmd := &cobra.Command{
		Use:   "add-allowlist-user",
		Short: "Allowlist a user from SAML lockdown allowing them to sign in using a password in addition to SAML.",
		Run: func(cmd *cobra.Command, args []string) {
			addAllowListUser(userId)
		},
	}
	cmd.Flags().StringVar(&userId, "userId", "", "Specify the id of the user")
	cmd.MarkFlagRequired("userId")
	return cmd
}

func addAllowListUser(userId string) {
	var allowListResponse api.GetSamlAllowListUsers
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/saml/allowlistedUsers/" + userId
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &allowListResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	allowListResponseJson, err := json.MarshalIndent(allowListResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(allowListResponseJson))
	}
}
