package created

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdPartitionCreate() *cobra.Command {
	var (
		name              string
		routingExpression string
		analyticsTier     string
		retentionPeriod   int
		isCompliant       bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new partition.",
		Run: func(cmd *cobra.Command, args []string) {
			createPartition(name, routingExpression, analyticsTier, retentionPeriod, isCompliant)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the partition")
	cmd.Flags().StringVar(&routingExpression, "routingExpression", "", "Specify the query that defines the data to be included in the partition")
	cmd.Flags().StringVar(&analyticsTier, "analyticsTier", "continuous", "Specify the Data Tier where the data in the partition will reside. "+
		"Possible values are continuous, frequent, infrequent.")
	cmd.Flags().IntVar(&retentionPeriod, "retentionPeriod", -1, "Specify the number of days to retain data in the partition. "+
		"-1 specifies that the default value for the account is used.")
	cmd.Flags().BoolVar(&isCompliant, "isCompliant", false, "Set to true if the partition is compliant")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("routingExpression")
	return cmd
}

func createPartition(name string, routingExpression string, analyticsTier string, retentionPeriod int, isCompliant bool) {
	var partitionsResponse api.Partitions
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreatePartition{
		Name:              name,
		RoutingExpression: routingExpression,
		AnalyticsTier:     analyticsTier,
		RetentionPeriod:   retentionPeriod,
		IsCompliant:       isCompliant,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "/v1/partitions"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
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
