package monitors

import (
	NewCmdMonitorsCreateFolder "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/monitors/create-folder"
	NewCmdMonitorsGetRootFolder "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/monitors/get-root-folder"
	NewCmdMonitorsGetUsageInfo "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/monitors/get-usage-info"
	"github.com/spf13/cobra"
)

func NewCmdMonitors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitors",
		Short: "Manage monitors",
		Long:  "Monitors continuously query your data to monitor and send notifications when specific events occur.",
	}
	cmd.AddCommand(NewCmdMonitorsCreateFolder.NewCmdMonitorsCreateFolder())
	cmd.AddCommand(NewCmdMonitorsGetRootFolder.NewCmdMonitorsGetRootFolder())
	cmd.AddCommand(NewCmdMonitorsGetUsageInfo.NewCmdMonitorsGetUsageInfo())
	return cmd
}
