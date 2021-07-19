package apps

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	NewCmdAppsGet "github.com/wizedkyle/sumocli/pkg/cmd/apps/get"
	NewCmdAppsInstall "github.com/wizedkyle/sumocli/pkg/cmd/apps/install"
	NewCmdAppsInstallStatus "github.com/wizedkyle/sumocli/pkg/cmd/apps/install_status"
	NewCmdAppsList "github.com/wizedkyle/sumocli/pkg/cmd/apps/list"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdApps(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apps",
		Short: "Manage apps (Beta)",
		Long:  "View and install Sumo Logic Applications that deliver out-of-the-box dashboards, saved searches, and field extraction for popular data sources.",
	}
	cmd.AddCommand(NewCmdAppsGet.NewCmdAppsGet(client, log))
	cmd.AddCommand(NewCmdAppsInstall.NewCmdAppsInstall(client, log))
	cmd.AddCommand(NewCmdAppsInstallStatus.NewCmdAppsInstallStatus(client, log))
	cmd.AddCommand(NewCmdAppsList.NewCmdAppsList(client, log))
	return cmd
}
