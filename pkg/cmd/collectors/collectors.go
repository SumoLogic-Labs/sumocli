package collectors

import (
	"github.com/spf13/cobra"
	cmdCollectorCreate "github.com/wizedkyle/sumocli/pkg/cmd/collectors/create"
	cmdCollectorDelete "github.com/wizedkyle/sumocli/pkg/cmd/collectors/delete"
	cmdCollectorGet "github.com/wizedkyle/sumocli/pkg/cmd/collectors/get"
	cmdCollectorsList "github.com/wizedkyle/sumocli/pkg/cmd/collectors/list"
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
	cmd.AddCommand(cmdCollectorDelete.NewCmdCollectorDelete())

	return cmd
}