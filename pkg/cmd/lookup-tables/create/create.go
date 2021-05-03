package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"strings"
)

func NewCmdLookupTablesCreate() *cobra.Command {
	var (
		description     string
		fieldNames      string
		fieldTypes      string
		primaryKeys     string
		ttl             int
		sizeLimitAction string
		name            string
		parentFolderId  string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new lookup table by providing a schema and specifying its configuration.",
		Run: func(cmd *cobra.Command, args []string) {
			createLookupTable(description, fieldNames, fieldTypes, primaryKeys, ttl, sizeLimitAction, name, parentFolderId)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the lookup table")
	cmd.Flags().StringVar(&fieldNames, "fieldNames", "", "List of field names (they need to be comma separated e.g. test,test1,test2")
	cmd.Flags().StringVar(&fieldTypes, "fieldTypes", "", "List of field types that align with the fieldNames "+
		"(they need to be comma separated e.g. string,boolean,int). The following fieldTypes can be specified: "+
		"boolean, int, long, double, string")
	cmd.Flags().StringVar(&primaryKeys, "primaryKeys", "", "List of field names that make up the primary key for the"+
		"lookup table (they need to be comma separated e.g. name1,name2,name3). ")
	cmd.Flags().IntVar(&ttl, "ttl", 0, "A time to live for each entry in the lookup table (in minutes). "+
		"365 days is the maximum ttl, leaving the ttl as 0 means that the records will not expire automatically.")
	cmd.Flags().StringVar(&sizeLimitAction, "sizeLimitAction", "StopIncomingMessages", "The action that needs to be taken"+
		"when the size limit is reached for the table. The possible values can be StopIncomingMessages (default) or DeleteOldData.")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the lookup table")
	cmd.Flags().StringVar(&parentFolderId, "parentFolderId", "", "Specify the parent folder path identifier of the lookup table in the Library")

	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("fieldNames")
	cmd.MarkFlagRequired("fieldTypes")
	cmd.MarkFlagRequired("primaryKeys")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("parentFolderId")
	return cmd
}

func createLookupTable(description string, fieldNames string, fieldTypes string, primaryKeys string,
	ttl int, sizeLimitAction string, name string, parentFolderId string) {
	var createLookupTableResponse api.LookupTableResponse
	log := logging.GetConsoleLogger()
	requestBodySchema := &api.CreateLookupTableRequest{
		Description:     description,
		PrimaryKeys:     strings.Split(primaryKeys, ","),
		Ttl:             ttl,
		SizeLimitAction: sizeLimitAction,
		Name:            name,
		ParentFolderId:  parentFolderId,
	}
	fieldNameSlice := strings.Split(fieldNames, ",")
	fieldTypeSlice := strings.Split(fieldTypes, ",")
	for i := range fieldNameSlice {
		fieldAddition := api.LookupTableFields{
			FieldName: fieldNameSlice[i],
			FieldType: fieldTypeSlice[i],
		}
		requestBodySchema.Fields = append(requestBodySchema.Fields, fieldAddition)
		i++
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/lookupTables"
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

	err = json.Unmarshal(responseBody, &createLookupTableResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	createLookupTableResponseJson, err := json.MarshalIndent(createLookupTableResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal lookupTableResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(createLookupTableResponseJson))
	}
}
