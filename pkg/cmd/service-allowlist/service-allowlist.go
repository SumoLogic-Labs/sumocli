package service_allowlist

import (
	"github.com/spf13/cobra"
	cmdServiceAllowlistAdd "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/add"
	cmdServiceAllowlistDisable "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/disable"
	cmdServiceAllowlistEnable "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/enable"
	cmdServiceAllowlistList "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/list"
	cmdServiceAllowlistRemove "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/remove"
	cmdServiceAllowlistStatus "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/status"
)

func NewCmdServiceAllowlist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service-allowlist",
		Short: "Manage the service allowlist",
		Long:  "Commands that all you to manage the Service Allowlist in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdServiceAllowlistAdd.NewCmdServiceAllowlistAdd())
	cmd.AddCommand(cmdServiceAllowlistDisable.NewCmdServiceAllowlistDisable())
	cmd.AddCommand(cmdServiceAllowlistEnable.NewCmdServiceAllowListEnable())
	cmd.AddCommand(cmdServiceAllowlistList.NewCmdServiceAllowlistList())
	cmd.AddCommand(cmdServiceAllowlistRemove.NewCmdServiceAllowlistRemove())
	cmd.AddCommand(cmdServiceAllowlistStatus.NewCmdServiceAllowlistStatus())
	return cmd
}
