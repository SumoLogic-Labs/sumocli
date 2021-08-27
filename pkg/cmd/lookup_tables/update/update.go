package update

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdLookupTablesEdit(client *cip.APIClient) *cobra.Command {
	var (
		description     string
		id              string
		ttl             int32
		sizeLimitAction string
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Edits the configuration of the given lookup table",
		Run: func(cmd *cobra.Command, args []string) {
			updateLookupTable(description, id, ttl, sizeLimitAction, client)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the lookup table")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table to edit")
	cmd.Flags().Int32Var(&ttl, "ttl", 0, "A time to live for each entry in the lookup table (in minutes). "+
		"365 days is the maximum ttl, leaving the ttl as 0 means that the records will not expire automatically.")
	cmd.Flags().StringVar(&sizeLimitAction, "sizeLimitAction", "StopIncomingMessages", "The action that needs to be taken"+
		"when the size limit is reached for the table. The possible values can be StopIncomingMessages (default) or DeleteOldData.")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updateLookupTable(description string, id string, ttl int32, sizeLimitAction string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.UpdateTable(types.LookupUpdateDefinition{
		Ttl:             ttl,
		Description:     description,
		SizeLimitAction: sizeLimitAction,
	},
		id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
