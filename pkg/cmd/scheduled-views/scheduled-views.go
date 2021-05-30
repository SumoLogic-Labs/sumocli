package scheduled_views

import (
	"github.com/spf13/cobra"
	NewCmdScheduledViewsCreate "github.com/wizedkyle/sumocli/pkg/cmd/scheduled-views/create"
	NewCmdScheduledViewsDisable "github.com/wizedkyle/sumocli/pkg/cmd/scheduled-views/disable"
	NewCmdScheduledViewsGet "github.com/wizedkyle/sumocli/pkg/cmd/scheduled-views/get"
	NewCmdScheduledViewsList "github.com/wizedkyle/sumocli/pkg/cmd/scheduled-views/list"
	NewCmdScheduledViewsPause "github.com/wizedkyle/sumocli/pkg/cmd/scheduled-views/pause"
	NewCmdScheduledViewsStart "github.com/wizedkyle/sumocli/pkg/cmd/scheduled-views/start"
	NewCmdScheduledViewsUpdate "github.com/wizedkyle/sumocli/pkg/cmd/scheduled-views/update"
)

func NewCmdScheduledViews() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scheduled-views",
		Short: "Managed scheduled views",
		Long:  "Scheduled Views speed the search process for small and historical subsets of your data by functioning as a pre-aggregated index.",
	}
	cmd.AddCommand(NewCmdScheduledViewsCreate.NewCmdScheduledViewsCreate())
	cmd.AddCommand(NewCmdScheduledViewsDisable.NewCmdScheduledViewsDisable())
	cmd.AddCommand(NewCmdScheduledViewsGet.NewCmdScheduledViewsGet())
	cmd.AddCommand(NewCmdScheduledViewsList.NewCmdScheduledViewsList())
	cmd.AddCommand(NewCmdScheduledViewsPause.NewCmdScheduledViewsPause())
	cmd.AddCommand(NewCmdScheduledViewsStart.NewCmdScheduledViewsStart())
	cmd.AddCommand(NewCmdScheduledViewsUpdate.NewCmdScheduledViewsUpdate())
	return cmd
}
