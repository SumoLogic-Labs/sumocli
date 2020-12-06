package delete

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/guard"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io/ioutil"
)

func NewCmdCollectorDelete() *cobra.Command {
	var (
		id			  string
	)
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic hosted collector",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Collector delete request started.")
			deleteCollector(id, logger)
			logger.Debug().Msg("Collector delete request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of the collector")

	return cmd
}

func deleteCollector(id string, logger zerolog.Logger) {
	var createCollectorResponse api.CollectorData
	guard.MustNotBeEmpty(id, "--id field needs to be specified.")

	client, request := factory.NewHttpRequest("DELETE", "v1/collectors/" + id)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		jsonErr := json.Unmarshal(responseBody, &createCollectorResponse)
		logging.LogError(jsonErr, logger)
		fmt.Println(createCollectorResponse.Name + " collector successfully deleted")
	}
}
