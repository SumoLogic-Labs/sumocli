package update

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
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
	data, response, err := client.UpdateTable(types.LookupUpdateDefinition{
		Ttl:             ttl,
		Description:     description,
		SizeLimitAction: sizeLimitAction,
	},
		id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
