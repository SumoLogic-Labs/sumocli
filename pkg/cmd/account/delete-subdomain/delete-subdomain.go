package delete_subdomain

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdAccountDeleteSubdomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-subdomain",
		Short: "Delete the configured subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteSubdomain()
		},
	}
	return cmd
}

func deleteSubdomain() {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/account/subdomain"
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
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
		fmt.Println("Subdomain was deleted.")
	}
}
