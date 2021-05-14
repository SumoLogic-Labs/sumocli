package apps

import (
	"github.com/spf13/cobra"
)

func NewCmdApps() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apps",
		Short: "Manage apps (Beta)",
		Long:  "View and install Sumo Logic Applications that deliver out-of-the-box dashboards, saved searches, and field extraction for popular data sources.",
	}
	return cmd
}
