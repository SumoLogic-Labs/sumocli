package upgrade

import (
	"github.com/spf13/cobra"
	cmdCollectorUpgradeBuilds "github.com/wizedkyle/sumocli/pkg/cmd/collectors/upgrade/get_available_builds"
	cmdCollectorUpgradeGet "github.com/wizedkyle/sumocli/pkg/cmd/collectors/upgrade/get_upgradable_collectors"
	cmdCollectorUpgradeStart "github.com/wizedkyle/sumocli/pkg/cmd/collectors/upgrade/start"
	cmdCollectorStatus "github.com/wizedkyle/sumocli/pkg/cmd/collectors/upgrade/status"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdUpgradeCollectors(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Manages the upgrading of collectors",
	}

	cmd.AddCommand(cmdCollectorUpgradeBuilds.NewCmdGetBuilds(client))
	cmd.AddCommand(cmdCollectorUpgradeGet.NewCmdGetUpgradableCollectors(client))
	cmd.AddCommand(cmdCollectorUpgradeStart.NewCmdUpgradeStart(client))
	cmd.AddCommand(cmdCollectorStatus.NewCmdUpgradableCollectorStatus(client))
	return cmd
}
