package delete_field

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldManagementDeleteField() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use: "delete-field",
		Short: "Deleting a field does not delete historical data assigned with that field. " +
			"If you delete a field by mistake and one or more of those dependencies break, you can re-add the field to get things working properly again. " +
			"You should always disable a field using sumocli field-management disable-custom-field and ensure things are behaving as expected before deleting a field.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteField(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteField(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/fields/" + id
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
		fmt.Println("The field was successfully deleted.")
	}
}
