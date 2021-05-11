package service_allowlist

import (
	"github.com/spf13/cobra"
	cmdServiceAllowlistList "github.com/wizedkyle/sumocli/pkg/cmd/service-allowlist/list"
)

func NewCmdServiceAllowlist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service-allowlist",
		Short: "Manage the service allowlist",
		Long:  "Commands that all you to manage the Service Allowlist in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdServiceAllowlistList.NewCmdServiceAllowlistList())
	return cmd
}
