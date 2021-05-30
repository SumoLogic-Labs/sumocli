package disable_custom_field

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldManagementDisableCustomField() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use: "disable-custom-field",
		Short: "After disabling a field Sumo Logic will start dropping its incoming values at ingest. " +
			"As a result, they won't be searchable or usable. Historical values are not removed and remain searchable.",
		Run: func(cmd *cobra.Command, args []string) {
			disableCustomField(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field")
	cmd.MarkFlagRequired("id")
	return cmd
}

func disableCustomField(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/fields/" + id + "/disable"
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
		fmt.Println("Field has been disabled.")
	}
}
