package remove

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/api"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmd/factory"
	"github.com/SumoLogic-Incubator/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdServiceAllowlistRemove() *cobra.Command {
	var (
		cidr        string
		description string
	)

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove allowlisted CIDR notations and/or IP addresses from the organization. Removed CIDRs/IPs will immediately lose access to Sumo Logic and content sharing.",
		Run: func(cmd *cobra.Command, args []string) {
			removeServiceAllowlistCidr(cidr, description)
		},
	}
	cmd.Flags().StringVar(&cidr, "cidr", "", "Specify the IP address to add in CIDR format")
	cmd.Flags().StringVar(&description, "description", "", "Specify the description for the IP address")
	cmd.MarkFlagRequired("cidr")
	return cmd
}

func removeServiceAllowlistCidr(cidr string, description string) {
	var allowlistCidrResponse api.ListServiceAllowlist
	log := logging.GetConsoleLogger()
	allowlistAddition := api.AllowlistCIDR{
		Cidr:        cidr,
		Description: description,
	}
	requestBodySchema := &api.ListServiceAllowlist{}
	requestBodySchema.Data = append(requestBodySchema.Data, allowlistAddition)
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "/v1/serviceAllowlist/addresses/remove"
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

	err = json.Unmarshal(responseBody, &allowlistCidrResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	allowlistCidrResponseJson, err := json.MarshalIndent(allowlistCidrResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(allowlistCidrResponseJson))
	}
}
