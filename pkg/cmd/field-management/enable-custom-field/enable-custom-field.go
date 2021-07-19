package enable_custom_field

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdFieldManagementEnableCustomField() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use: "enable-custom-field",
		Short: "Fields have to be enabled to be assigned to your data. " +
			"This operation ensures that a specified field is enabled and Sumo Logic will treat it as safe to process. " +
			"All created custom fields using sumocli field-management create-field are enabled by default.",
		Run: func(cmd *cobra.Command, args []string) {
			enableCustomField(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field")
	cmd.MarkFlagRequired("id")
	return cmd
}

func enableCustomField(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/fields/" + id + "/enable"
	client, request := factory.NewHttpRequest("PUT", requestUrl)
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
		fmt.Println("Field has been enabled.")
	}
}
