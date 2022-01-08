package delete

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdDashboardsDelete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a dashboard by the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteDashboards(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the dashboard")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteDashboards(id string, client *cip.APIClient) {
	response, err := client.DeleteDashboard(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Dashboard was deleted successfully.")
	}
}
