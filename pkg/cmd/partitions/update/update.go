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

func NewCmdPartitionUpdate() *cobra.Command {
	var (
		id                               string
		retentionPeriod                  int
		reduceRetentionPeriodImmediately bool
		isCompliant                      bool
		routingExpression                string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing partition in the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			updatePartition(id, retentionPeriod, reduceRetentionPeriodImmediately, isCompliant, routingExpression)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the partition")
	cmd.Flags().IntVar(&retentionPeriod, "retentionPeriod", -1, "Specify the number of days to retain data in the partition. "+
		"-1 specifies that the default value for the account is used.")
	cmd.Flags().BoolVar(&reduceRetentionPeriodImmediately, "reduceRetentionPeriodImmediately", false, "This is required if the newly specified retentionPeriod is less than the existing retention period. "+
		"A value of true says that data between the existing retention period and the new retention period should be deleted immediately. "+
		"If false, such data will be deleted after seven days.")
	cmd.Flags().BoolVar(&isCompliant, "isCompliant", false, "Set to true if the partition is compliant")
	cmd.Flags().StringVar(&routingExpression, "routingExpression", "", "Specify the query that defines the data to be included in the partition")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updatePartition(id string, retentionPeriod int, reduceRetentionPeriodImmediately bool, isCompliant bool,
	routingExpression string) {
	var partitionsResponse api.Partitions
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.UpdatePartition{
		RetentionPeriod:                  retentionPeriod,
		ReduceRetentionPeriodImmediately: reduceRetentionPeriodImmediately,
		IsCompliant:                      isCompliant,
		RoutingExpression:                routingExpression,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "/v1/partitions/" + id
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

	err = json.Unmarshal(responseBody, &partitionsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	partitionsResponseJson, err := json.MarshalIndent(partitionsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(partitionsResponseJson))
	}
}
