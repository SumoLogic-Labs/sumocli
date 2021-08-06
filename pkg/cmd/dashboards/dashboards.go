package dashboards

import (
	"github.com/spf13/cobra"
	NewCmdDashboardCreate "github.com/wizedkyle/sumocli/pkg/cmd/dashboards/create"
	NewCmdDashboardDelete "github.com/wizedkyle/sumocli/pkg/cmd/dashboards/delete"
	NewCmdDashboardsGet "github.com/wizedkyle/sumocli/pkg/cmd/dashboards/get"
	NewCmdDashboardsUpdate "github.com/wizedkyle/sumocli/pkg/cmd/dashboards/update"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdDashboards(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dashboards",
		Short: "Manage dashboards (New)",
		Long:  "Commands that allow you to create, modify or delete new dashboards.",
	}
	cmd.AddCommand(NewCmdDashboardCreate.NewCmdDashboardsCreate())
	cmd.AddCommand(NewCmdDashboardDelete.NewCmdDashboardsDelete(client))
	cmd.AddCommand(NewCmdDashboardsGet.NewCmdDashboardsGet(client))
	cmd.AddCommand(NewCmdDashboardsUpdate.NewCmdDashboardsUpdate())
	return cmd
}
