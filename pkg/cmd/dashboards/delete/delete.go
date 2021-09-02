package delete

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
	httpResponse, errorResponse := client.DeleteDashboard(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Dashboard was deleted successfully.")
	}
}
