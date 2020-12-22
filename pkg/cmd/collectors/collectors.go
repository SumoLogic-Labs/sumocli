package collectors

import (
	"github.com/spf13/cobra"
	cmdCollectorCreate "github.com/wizedkyle/sumocli/pkg/cmd/collectors/create"
	cmdCollectorList "github.com/wizedkyle/sumocli/pkg/cmd/collectors/list"
)

func NewCmdCollectors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collectors <command>",
		Short: "Managed collectors",
	}

	cmd.AddCommand(cmdCollectorCreate.NewCmdCollectorCreate())
	cmd.AddCommand(cmdCollectorList.NewCmdCollectorList())
	return cmd
}
