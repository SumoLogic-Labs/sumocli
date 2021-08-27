package create

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdLookupTablesCreate(client *cip.APIClient) *cobra.Command {
	var (
		description     string
		fieldNames      []string
		fieldTypes      []string
		primaryKeys     []string
		ttl             int32
		sizeLimitAction string
		name            string
		parentFolderId  string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new lookup table by providing a schema and specifying its configuration.",
		Run: func(cmd *cobra.Command, args []string) {
			createLookupTable(description, fieldNames, fieldTypes, primaryKeys, ttl, sizeLimitAction, name,
				parentFolderId, client)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the lookup table")
	cmd.Flags().StringSliceVar(&fieldNames, "fieldNames", []string{}, "List of field names (they need to be comma separated e.g. test,test1,test2")
	cmd.Flags().StringSliceVar(&fieldTypes, "fieldTypes", []string{}, "List of field types that align with the fieldNames "+
		"(they need to be comma separated e.g. string,boolean,int). The following fieldTypes can be specified: "+
		"boolean, int, long, double, string")
	cmd.Flags().StringSliceVar(&primaryKeys, "primaryKeys", []string{}, "List of field names that make up the primary key for the"+
		"lookup table (they need to be comma separated e.g. name1,name2,name3). ")
	cmd.Flags().Int32Var(&ttl, "ttl", 0, "A time to live for each entry in the lookup table (in minutes). "+
		"365 days is the maximum ttl, leaving the ttl as 0 means that the records will not expire automatically.")
	cmd.Flags().StringVar(&sizeLimitAction, "sizeLimitAction", "StopIncomingMessages", "The action that needs to be taken "+
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

func createLookupTable(description string, fieldNames []string, fieldTypes []string, primaryKeys []string,
	ttl int32, sizeLimitAction string, name string, parentFolderId string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.CreateTable(types.LookupTableDefinition{
		Description:     description,
		Fields:          cmdutils.GenerateLookupTableFields(fieldNames, fieldTypes),
		PrimaryKeys:     primaryKeys,
		Ttl:             ttl,
		SizeLimitAction: sizeLimitAction,
		Name:            name,
		ParentFolderId:  parentFolderId,
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
