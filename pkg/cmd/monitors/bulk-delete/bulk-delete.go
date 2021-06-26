package bulk_delete

import (
	"github.com/spf13/cobra"
)

func NewCmdMonitorsBulkDelete() *cobra.Command {
	var ids []string

	cmd := &cobra.Command{
		Use:   "bulk-delete",
		Short: "Bulk delete a monitor or folder by the given identifiers in the monitors library.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.Flags().StringSliceVar(&ids, "ids", []string{}, "Specify the ids to delete (comma separated)")
	cmd.MarkFlagRequired("ids")
	return cmd
}

func bulkDeleteMonitors(ids []string) {

}
