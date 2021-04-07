package delete

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdUserDelete() *cobra.Command {
	var (
		id         string
		transferTo string
	)

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic user",
		Run: func(cmd *cobra.Command, args []string) {
			deleteUser(id, transferTo)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user to delete")
	cmd.Flags().StringVar(&transferTo, "transferTo", "", "Specify the id of the user to transfer data to")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteUser(id string, transferTo string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/users/" + id
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		err := json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("error unmarshalling response body")
		}
	} else {
		fmt.Println("User was deleted.")
	}
}
