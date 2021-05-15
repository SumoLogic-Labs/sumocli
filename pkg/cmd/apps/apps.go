package apps

import (
	"github.com/spf13/cobra"
	NewCmdAppsGet "github.com/wizedkyle/sumocli/pkg/cmd/apps/get"
	NewCmdAppsInstall "github.com/wizedkyle/sumocli/pkg/cmd/apps/install"
	NewCmdAppsInstallStatus "github.com/wizedkyle/sumocli/pkg/cmd/apps/install-status"
	NewCmdAppsList "github.com/wizedkyle/sumocli/pkg/cmd/apps/list"
)

func NewCmdApps() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apps",
		Short: "Manage apps (Beta)",
		Long:  "View and install Sumo Logic Applications that deliver out-of-the-box dashboards, saved searches, and field extraction for popular data sources.",
	}
	cmd.AddCommand(NewCmdAppsGet.NewCmdAppsGet())
	cmd.AddCommand(NewCmdAppsInstall.NewCmdAppsInstall())
	cmd.AddCommand(NewCmdAppsInstallStatus.NewCmdAppsInstallStatus())
	cmd.AddCommand(NewCmdAppsList.NewCmdAppsList())
	return cmd
}
