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

func NewCmdUserList() *cobra.Command {
	var (
		email           string
		numberOfResults string
		sortBy          string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic users",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User list request started.")
			listUsers(email, numberOfResults, sortBy, logger)
			logger.Debug().Msg("User list request finished.")
		},
	}

	cmd.Flags().StringVar(&email, "email", "", "Specify the email address of the user")
	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&sortBy, "sort", "", "Sort the results by firstName, lastName or email")

	return cmd
}

func listUsers(email string, numberOfResults string, sortBy string, logger zerolog.Logger) {
	var userInfo api.Users

	client, request := factory.NewHttpRequest("GET", "v1/users")
	query := url.Values{}
	if numberOfResults != "" {
		query.Add("limit", numberOfResults)
	}
	if sortBy != "" {
		if factory.ValidateUserSortBy(sortBy) == false {
			fmt.Println(sortBy + "is an invalid field to sort by. Available fields are firstName, lastName or email. ")
		} else {
			query.Add("sortBy", sortBy)
		}
	}
	if email != "" {
		query.Add("email", email)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	logging.LogError(err, logger)

	jsonErr := json.Unmarshal(responseBody, &userInfo)
	logging.LogError(jsonErr, logger)

	userInfoJson, err := json.MarshalIndent(userInfo.Data, "", "    ")
	logging.LogError(err, logger)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		fmt.Println(string(userInfoJson))
	}
}
