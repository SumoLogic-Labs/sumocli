package create

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

func NewCmdCollectorCreate() *cobra.Command {
	var (
		name          string
		description   string
		category      string
		hostName      string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic hosted collector",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Collector create request started.")
			createCollector(name, description, "Hosted", category, hostName, logger)
			logger.Debug().Msg("Collector create request finished.")
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Name of the collector")
	cmd.Flags().StringVar(&description, "description", "", "Description for the collector")
	cmd.Flags().StringVar(&category, "category", "", "Category of the collector")
	cmd.Flags().StringVar(&hostName, "hostName", "", "HostName of the collector")

	return cmd
}

func createCollector(name string, description string, collectorType string, category string, hostName string, logger zerolog.Logger) {
	var createCollectorResponse api.CollectorData
	guard.MustNotBeEmpty(name, "--name field needs to be specified.")
	guard.MustNotBeEmpty(description, "--description field needs to be specified.")

	requestBodySchema := &api.CreateCollectorRequest{
		Collector: api.CollectorCreateData{
			Name: name,
			Description: description,
			CollectorType: collectorType,
			Category: category,
			HostName: hostName,
		},
	}

	requestBody, _ := json.Marshal(requestBodySchema)
	client, request := factory.NewHttpRequestWithBody("POST", "v1/collectors", requestBody)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		jsonErr := json.Unmarshal(responseBody, &createCollectorResponse)
		logging.LogError(jsonErr, logger)
		fmt.Println(createCollectorResponse.Name + " collector successfully created")
	}
}
