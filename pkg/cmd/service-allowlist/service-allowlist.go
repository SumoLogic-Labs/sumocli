package service_allowlist

import (
	"github.com/spf13/cobra"
	cmdServiceAllowListDisable "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/disable"
	cmdServiceAllowlistEnable "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/enable"
	cmdServiceAllowlistList "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/list"
	cmdServiceAllowlistStatus "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/status"
)

func NewCmdServiceAllowlist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service-allowlist",
		Short: "Manage the service allowlist",
		Long:  "Commands that all you to manage the Service Allowlist in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdServiceAllowListDisable.NewCmdServiceAllowlistDisable())
	cmd.AddCommand(cmdServiceAllowlistEnable.NewCmdServiceAllowListEnable())
	cmd.AddCommand(cmdServiceAllowlistList.NewCmdServiceAllowlistList())
	cmd.AddCommand(cmdServiceAllowlistStatus.NewCmdServiceAllowlistStatus())
	return cmd
}
