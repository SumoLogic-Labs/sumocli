package disable

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

func NewCmdServiceAllowlistDisable() *cobra.Command {
	var (
		login   bool
		content bool
		both    bool
	)

	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disable service allowlisting functionality for login/API authentication or content sharing for the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			disableServiceAllowlist(login, content, both)
		},
	}
	cmd.Flags().BoolVar(&login, "login", false, "Set to true if you want the allowlist to affect logins")
	cmd.Flags().BoolVar(&content, "content", false, "Set to true if you want the allowlist to affect content")
	cmd.Flags().BoolVar(&both, "both", false, "Set to true if you want the allowlist to affect both logins and content")
	return cmd
}

func disableServiceAllowlist(login bool, content bool, both bool) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/serviceAllowlist/disable"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	query := url.Values{}
	if login == true {
		query.Add("allowlistType", "Login")
	} else if content == true {
		query.Add("allowlistType", "Content")
	} else if both == true {
		query.Add("allowlistType", "Both")
	} else {
		log.Fatal().Msg("Please set either login, content or both to true")
	}
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
		fmt.Println("Service Allowlist is disabled.")
	}
}
