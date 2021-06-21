package monitors

import (
	"github.com/spf13/cobra"
	NewCmdMonitorsCreateFolder "github.com/wizedkyle/sumocli/pkg/cmd/monitors/create-folder"
	NewCmdMonitorsGetRootFolder "github.com/wizedkyle/sumocli/pkg/cmd/monitors/get-root-folder"
	NewCmdMonitorsGetUsageInfo "github.com/wizedkyle/sumocli/pkg/cmd/monitors/get-usage-info"
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
