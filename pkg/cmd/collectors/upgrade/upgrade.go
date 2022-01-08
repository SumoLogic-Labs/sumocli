package upgrade

import (
	cmdCollectorUpgradeBuilds "github.com/SumoLogic-Labs/sumocli/pkg/cmd/collectors/upgrade/get_available_builds"
	cmdCollectorUpgradeGet "github.com/SumoLogic-Labs/sumocli/pkg/cmd/collectors/upgrade/get_upgradable_collectors"
	cmdCollectorUpgradeStart "github.com/SumoLogic-Labs/sumocli/pkg/cmd/collectors/upgrade/start"
	cmdCollectorStatus "github.com/SumoLogic-Labs/sumocli/pkg/cmd/collectors/upgrade/status"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
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
