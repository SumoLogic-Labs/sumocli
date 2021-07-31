package apps

import (
	"github.com/spf13/cobra"
	NewCmdAppsGet "github.com/wizedkyle/sumocli/pkg/cmd/apps/get"
	NewCmdAppsInstall "github.com/wizedkyle/sumocli/pkg/cmd/apps/install"
	NewCmdAppsInstallStatus "github.com/wizedkyle/sumocli/pkg/cmd/apps/install_status"
	NewCmdAppsList "github.com/wizedkyle/sumocli/pkg/cmd/apps/list"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdApps(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apps",
		Short: "Manage apps (Beta)",
		Long:  "View and install Sumo Logic Applications that deliver out-of-the-box dashboards, saved searches, and field extraction for popular data sources.",
	}
	cmd.AddCommand(NewCmdAppsGet.NewCmdAppsGet(client))
	cmd.AddCommand(NewCmdAppsInstall.NewCmdAppsInstall(client))
	cmd.AddCommand(NewCmdAppsInstallStatus.NewCmdAppsInstallStatus(client))
	cmd.AddCommand(NewCmdAppsList.NewCmdAppsList(client))
	return cmd
}
