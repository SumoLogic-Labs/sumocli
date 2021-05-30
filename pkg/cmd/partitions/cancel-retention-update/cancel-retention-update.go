package cancel_retention_update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdPartitionsCancelRetentionUpdate() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "cancel-retention-update",
		Short: "Cancel update to retention of a partition for which retention was updated previously using sumocli partitions update and the reduceRetentionPeriodImmediately parameter was set to false",
		Run: func(cmd *cobra.Command, args []string) {
			cancelRetentionUpdate(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the partition")
	cmd.MarkFlagRequired("id")
	return cmd
}

func cancelRetentionUpdate(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/partitions/" + id + "/cancelRetentionUpdate"
	client, request := factory.NewHttpRequest("POST", requestUrl)
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
		fmt.Println("The retention update was cancelled successfully.")
	}
}
