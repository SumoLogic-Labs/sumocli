package recover_subdomain

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"net/url"
)

func NewCmdAccountRecoverSubdomain() *cobra.Command {
	var email string

	cmd := &cobra.Command{
		Use:   "recover-subdomain",
		Short: "Send an email with the subdomain information for a user with the given email address.",
		Run: func(cmd *cobra.Command, args []string) {
			recoverSubdomain(email)
		},
	}
	cmd.Flags().StringVar(&email, "email", "", "Specify an email address of the user to get subdomain information")
	cmd.MarkFlagRequired("email")
	return cmd
}

func recoverSubdomain(email string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/account/subdomain/recover"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	query := url.Values{}
	query.Add("email", email)
	request.URL.RawQuery = query.Encode()
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
		err := json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
	} else {
		fmt.Println("An email containing information about associated subdomains for the given email was sent.")
	}
}
