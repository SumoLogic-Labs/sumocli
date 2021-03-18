package list

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"net/url"
)

func NewCmdRoleList() *cobra.Command {
	var (
		numberOfResults string
		filter          string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic roles",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role list request started.")
			listRoles(numberOfResults, filter, logger)
			logger.Debug().Msg("Role list request finished.")
		},
	}

	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")

	return cmd
}

func listRoles(numberOfResults string, filter string, logger zerolog.Logger) {
	var roleInfo api.Role

	client, request := factory.NewHttpRequest("GET", "v1/roles")
	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if filter != "" {
		query.Add("name", filter)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	logging.LogError(err, logger)

	jsonErr := json.Unmarshal(responseBody, &roleInfo)
	logging.LogError(jsonErr, logger)

	roleInfoJson, err := json.MarshalIndent(roleInfo.Data, "", "    ")
	logging.LogError(err, logger)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		fmt.Println(string(roleInfoJson))
	}
}
