package edit

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdLookupTablesEdit() *cobra.Command {
	var (
		description     string
		id              string
		ttl             int
		sizeLimitAction string
	)

	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edits the configuration of the given lookup table",
		Run: func(cmd *cobra.Command, args []string) {
			editLookupTable(description, id, ttl, sizeLimitAction)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the lookup table")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table to edit")
	cmd.Flags().IntVar(&ttl, "ttl", 0, "A time to live for each entry in the lookup table (in minutes). "+
		"365 days is the maximum ttl, leaving the ttl as 0 means that the records will not expire automatically.")
	cmd.Flags().StringVar(&sizeLimitAction, "sizeLimitAction", "StopIncomingMessages", "The action that needs to be taken"+
		"when the size limit is reached for the table. The possible values can be StopIncomingMessages (default) or DeleteOldData.")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("id")
	return cmd
}

func editLookupTable(description string, id string, ttl int, sizeLimitAction string) {
	var editLookupTableResponse api.LookupTableResponse
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.EditLookupTable{
		Description:     description,
		Ttl:             ttl,
		SizeLimitAction: sizeLimitAction,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/lookupTables/" + id
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

	err = json.Unmarshal(responseBody, &editLookupTableResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	editLookupTableResponseJson, err := json.MarshalIndent(editLookupTableResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(editLookupTableResponseJson))
	}
}
