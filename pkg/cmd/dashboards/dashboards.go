package dashboards

import (
	NewCmdDashboardCreate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/dashboards/create"
	NewCmdDashboardDelete "github.com/SumoLogic-Labs/sumocli/pkg/cmd/dashboards/delete"
	NewCmdDashboardsGet "github.com/SumoLogic-Labs/sumocli/pkg/cmd/dashboards/get"
	NewCmdDashboardsUpdate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/dashboards/update"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdDashboards(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dashboards",
		Short: "Manage dashboards (New)",
		Long:  "Commands that allow you to create, modify or delete new dashboards.",
	}
	cmd.AddCommand(NewCmdDashboardCreate.NewCmdDashboardsCreate(client))
	cmd.AddCommand(NewCmdDashboardDelete.NewCmdDashboardsDelete(client))
	cmd.AddCommand(NewCmdDashboardsGet.NewCmdDashboardsGet(client))
	cmd.AddCommand(NewCmdDashboardsUpdate.NewCmdDashboardsUpdate(client))
	return cmd
}
