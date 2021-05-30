package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdScheduledViewsUpdate() *cobra.Command {
	var (
		id                               string
		dataForwardingId                 string
		retentionPeriod                  int
		reduceRetentionPeriodImmediately bool
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing scheduled view.",
		Run: func(cmd *cobra.Command, args []string) {
			updateScheduledView(id, dataForwardingId, retentionPeriod, reduceRetentionPeriodImmediately)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the scheduled view")
	cmd.Flags().StringVar(&dataForwardingId, "dataForwardingId", "", "Specify an ID of a data forwarding configuration to be used by the scheduled view")
	cmd.Flags().IntVar(&retentionPeriod, "retentionPeriod", -1, "Specify the number of days to retain data in the partition. "+
		"-1 specifies that the default value for the account is used.")
	cmd.Flags().BoolVar(&reduceRetentionPeriodImmediately, "reduceRetentionPeriodImmediately", false, "This is required if the newly specified retentionPeriod is less than the existing retention period. "+
		"A value of true says that data between the existing retention period and the new retention period should be deleted immediately. "+
		"If false, such data will be deleted after seven days.")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updateScheduledView(id string, dataForwardingId string, retentionPeriod int, reduceRetentionPeriodImmediately bool) {
	var scheduledViewResponse api.ScheduledViews
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.UpdateScheduledView{
		DataForwardingId:                 dataForwardingId,
		RetentionPeriod:                  retentionPeriod,
		ReduceRetentionPeriodImmediately: reduceRetentionPeriodImmediately,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/scheduledViews/" + id
	client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &scheduledViewResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	scheduledViewResponseJson, err := json.MarshalIndent(scheduledViewResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(scheduledViewResponseJson))
	}
}
