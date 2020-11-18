package collectors

import (
	"github.com/spf13/cobra"
	cmdCollectorsList "github.com/wizedkyle/sumocli/pkg/cmd/collectors/list"
	cmdCollectorGet "github.com/wizedkyle/sumocli/pkg/cmd/collectors/get"
	cmdCollectorCreate "github.com/wizedkyle/sumocli/pkg/cmd/collectors/create"
)

func NewCmdCollector() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collectors <command>",
		Short: "Manage collectors",
		Long:  "Work with Sumo Logic collectors",
	}

	cmd.AddCommand(cmdCollectorsList.NewCmdCollectorsList())
	cmd.AddCommand(cmdCollectorGet.NewCmdCollectorGet())
	cmd.AddCommand(cmdCollectorCreate.NewCmdCollectorCreate())

	return cmd
}