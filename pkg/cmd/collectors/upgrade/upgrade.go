package upgrade

import (
	"github.com/spf13/cobra"
	cmdCollectorUpgradeBuilds "github.com/wizedkyle/sumocli/pkg/cmd/collectors/upgrade/builds"
	cmdCollectorUpgradeGet "github.com/wizedkyle/sumocli/pkg/cmd/collectors/upgrade/get"
)

func NewCmdUpgradeCollectors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade <command>",
		Short: "Manages the upgrading of collectors",
	}

	cmd.AddCommand(cmdCollectorUpgradeBuilds.NewCmdGetBuilds())
	cmd.AddCommand(cmdCollectorUpgradeGet.NewCmdGetUpgradableCollectors())
	return cmd
}
