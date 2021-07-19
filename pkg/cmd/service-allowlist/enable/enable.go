package enable

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

func NewCmdServiceAllowListEnable() *cobra.Command {
	var (
		login   bool
		content bool
		both    bool
	)

	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable service allowlisting functionality for the organization.",
		Long: "The service allowlisting can be for 1. Login: If enabled, access to Sumo Logic is granted only to CIDRs/IP addresses that are allowlisted. " +
			"2. Content: If enabled, dashboards can be shared with users connecting from CIDRs/IP addresses that are allowlisted without logging in.",
		Run: func(cmd *cobra.Command, args []string) {
			enableServiceAllowlist(login, content, both)
		},
	}
	cmd.Flags().BoolVar(&login, "login", false, "Set to true if you want the allowlist to affect logins")
	cmd.Flags().BoolVar(&content, "content", false, "Set to true if you want the allowlist to affect content")
	cmd.Flags().BoolVar(&both, "both", false, "Set to true if you want the allowlist to affect both logins and content")
	return cmd
}

func enableServiceAllowlist(login bool, content bool, both bool) {
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/serviceAllowlist/enable"
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
		fmt.Println("Service Allowlist is enabled.")
	}
}
