package create

import "github.com/spf13/cobra"

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

		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the lookup table")
	cmd.Flags().StringVar(&fieldNames, "fieldNames", "", "List of field names (they need to be comma separated e.g. test,test1,test2")
	cmd.Flags().StringVar(&fieldTypes, "fieldTypes", "", "List of field types that align with the fieldNames "+
		"(they need to be comma separated e.g. string,boolean,int). The following fieldTypes can be specified: "+
		"boolean, int, long, double, string")

	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("fieldNames")
	cmd.MarkFlagRequired("fieldTypes")
	return cmd
}

func createLookupTable() {

}
