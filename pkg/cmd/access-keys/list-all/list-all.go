package list_all

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"net/url"
	"strconv"
)

func NewCmdAccessKeysListAll() *cobra.Command {
	var limit int

	cmd := &cobra.Command{
		Use:   "list-all",
		Short: "List all access keys in your account.",
		Run: func(cmd *cobra.Command, args []string) {
			listAllAccessKeys(limit)
		},
	}
	cmd.Flags().IntVar(&limit, "limit", 100, "Specify the number of access keys returned")
	return cmd
}

func listAllAccessKeys(limit int) {
	var accessKeyResponse api.ListAccessKeysResponse
	log := logging.GetConsoleLogger()
	requestUrl := "v1/accessKeys"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &accessKeyResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	accessKeyResponseJson, err := json.MarshalIndent(accessKeyResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(accessKeyResponseJson))
	}
}
